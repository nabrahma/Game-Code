import React, { useState } from 'react';
import { ArrowUp } from 'lucide-react';
import { clsx } from 'clsx';

interface UpvoteButtonProps {
  initialUpvotes: number;
  initialHasUpvoted?: boolean;
  onToggle: () => Promise<boolean>; // Returns true if added, false if removed
}

export function UpvoteButton({ initialUpvotes, initialHasUpvoted = false, onToggle }: UpvoteButtonProps) {
  const [upvotes, setUpvotes] = useState(initialUpvotes);
  const [hasUpvoted, setHasUpvoted] = useState(initialHasUpvoted);
  const [loading, setLoading] = useState(false);

  const handleClick = async () => {
    if (loading) return;
    setLoading(true);
    
    // Optimistic UI update
    const previousUpvoted = hasUpvoted;
    const previousCount = upvotes;
    
    setHasUpvoted(!previousUpvoted);
    setUpvotes(prev => previousUpvoted ? prev - 1 : prev + 1);

    try {
      const added = await onToggle();
      // Ensure sync with server state
      setHasUpvoted(added);
      if (added !== !previousUpvoted) {
        setUpvotes(previousCount + (added ? 1 : 0));
      }
    } catch (err) {
      // Revert on error
      setHasUpvoted(previousUpvoted);
      setUpvotes(previousCount);
    } finally {
      setLoading(false);
    }
  };

  return (
    <button
      onClick={handleClick}
      disabled={loading}
      className={clsx(
        "flex items-center gap-1.5 px-2.5 py-1 rounded-md text-sm font-medium transition-colors border",
        hasUpvoted 
          ? "bg-indigo-500/10 text-indigo-400 border-indigo-500/30 hover:bg-indigo-500/20"
          : "bg-zinc-800/50 text-zinc-400 border-zinc-700 hover:bg-zinc-800 hover:text-zinc-200"
      )}
    >
      <ArrowUp className={clsx("h-4 w-4", hasUpvoted && "stroke-2")} />
      <span>{upvotes}</span>
    </button>
  );
}
