import React, { useState, useEffect } from "react";
import { Plus } from "lucide-react";

interface ProblemList {
  id: string;
  title: string;
  slug: string;
}

export function AddToListDropdown({ problemId }: { problemId: string }) {
  const [isOpen, setIsOpen] = useState(false);
  const [lists, setLists] = useState<ProblemList[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (isOpen && lists.length === 0) {
      setLoading(true);
      // Fetch user's lists
      fetch('http://localhost:8080/api/lists/user?user_id=00000000-0000-0000-0000-000000000001')
        .then(res => res.json())
        .then(data => setLists(data || []))
        .catch(err => console.error(err))
        .finally(() => setLoading(false));
    }
  }, [isOpen, lists.length]);

  const addToList = async (listId: string) => {
    try {
      const res = await fetch(`http://localhost:8080/api/lists/${listId}/problems`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ problem_id: problemId })
      });
      if (res.ok) {
        setIsOpen(false);
        // Optional: show a toast success message
      }
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="relative">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="flex h-8 w-8 items-center justify-center rounded-md text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-white"
        title="Add to List"
      >
        <Plus className="h-5 w-5" />
      </button>

      {isOpen && (
        <div className="absolute right-0 mt-2 w-48 rounded-md border border-zinc-800 bg-zinc-900 shadow-lg z-50">
          <div className="p-2 text-xs font-semibold text-zinc-500 uppercase tracking-wider">
            Save to List
          </div>
          <div className="max-h-48 overflow-y-auto">
            {loading ? (
              <div className="p-2 text-sm text-zinc-400">Loading...</div>
            ) : lists.length === 0 ? (
              <div className="p-2 text-sm text-zinc-400">No lists found</div>
            ) : (
              lists.map((list) => (
                <button
                  key={list.id}
                  onClick={() => addToList(list.id)}
                  className="w-full text-left px-3 py-2 text-sm text-zinc-300 hover:bg-emerald-500/20 hover:text-emerald-400 transition-colors"
                >
                  {list.title}
                </button>
              ))
            )}
          </div>
        </div>
      )}
    </div>
  );
}
