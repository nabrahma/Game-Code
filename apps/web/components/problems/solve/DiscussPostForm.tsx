import React, { useState } from 'react';
import { Send } from 'lucide-react';

interface DiscussPostFormProps {
  slug: string;
  onSuccess: () => void;
}

export function DiscussPostForm({ slug, onSuccess }: DiscussPostFormProps) {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim() || !content.trim()) return;

    setLoading(true);
    setError('');

    try {
      const res = await fetch(`http://localhost:8080/api/problems/${slug}/discuss`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, content }),
      });

      if (!res.ok) throw new Error('Failed to post discussion');
      
      setTitle('');
      setContent('');
      onSuccess();
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-zinc-900 border border-zinc-800 rounded-xl p-4 mb-8">
      <h3 className="text-lg font-semibold text-white mb-4">Start a Discussion</h3>
      
      {error && <div className="bg-red-500/10 text-red-400 p-3 rounded-lg mb-4 text-sm">{error}</div>}
      
      <div className="space-y-4">
        <div>
          <input
            type="text"
            placeholder="What's on your mind? (e.g., How to optimize the loop?)"
            className="w-full bg-zinc-950 border border-zinc-800 rounded-lg px-4 py-2 text-white placeholder-zinc-500 focus:outline-none focus:border-indigo-500 transition-colors"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            maxLength={100}
          />
        </div>
        <div>
          <textarea
            placeholder="Share your thoughts, ask a question, or post a hint..."
            className="w-full bg-zinc-950 border border-zinc-800 rounded-lg px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-indigo-500 transition-colors h-32 resize-y"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
          />
        </div>
        <div className="flex justify-end">
          <button
            type="submit"
            disabled={loading || !title.trim() || !content.trim()}
            className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 disabled:bg-zinc-800 disabled:text-zinc-500 text-white px-4 py-2 rounded-lg font-medium transition-colors"
          >
            {loading ? 'Posting...' : (
              <>
                <Send className="w-4 h-4" />
                Post
              </>
            )}
          </button>
        </div>
      </div>
    </form>
  );
}
