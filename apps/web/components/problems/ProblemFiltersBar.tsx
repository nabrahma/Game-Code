import React from "react";
import { Search } from "lucide-react";

interface ProblemFiltersBarProps {
  searchQuery: string;
  setSearchQuery: (q: string) => void;
  difficulty: string;
  setDifficulty: (d: string) => void;
}

export function ProblemFiltersBar({
  searchQuery,
  setSearchQuery,
  difficulty,
  setDifficulty,
}: ProblemFiltersBarProps) {
  return (
    <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between mb-6">
      <div className="relative max-w-md flex-1">
        <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
          <Search className="h-4 w-4 text-zinc-500" />
        </div>
        <input
          type="text"
          className="block w-full rounded-md border border-zinc-800 bg-zinc-900 py-2 pl-10 pr-3 text-sm text-zinc-200 placeholder-zinc-500 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
          placeholder="Search problems..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
        />
      </div>

      <div className="flex items-center gap-2">
        <select
          value={difficulty}
          onChange={(e) => setDifficulty(e.target.value)}
          className="rounded-md border border-zinc-800 bg-zinc-900 py-2 pl-3 pr-8 text-sm text-zinc-300 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
        >
          <option value="">All Difficulties</option>
          <option value="easy">Easy</option>
          <option value="medium">Medium</option>
          <option value="hard">Hard</option>
        </select>
      </div>
    </div>
  );
}
