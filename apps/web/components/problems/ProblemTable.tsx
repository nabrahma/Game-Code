import React from "react";
import Link from "next/link";
import { ProblemSummary } from "@/lib/hooks/useProblems";
import { DifficultyBadge } from "../ui/DifficultyBadge";

interface ProblemTableProps {
  problems: ProblemSummary[];
  isLoading: boolean;
}

export function ProblemTable({ problems, isLoading }: ProblemTableProps) {
  if (isLoading) {
    return (
      <div className="flex justify-center p-8 text-zinc-500">
        <div className="h-6 w-6 animate-spin rounded-full border-2 border-indigo-500 border-t-transparent" />
      </div>
    );
  }

  if (problems.length === 0) {
    return (
      <div className="flex justify-center p-8 text-zinc-500">
        No problems found matching your criteria.
      </div>
    );
  }

  return (
    <div className="overflow-hidden rounded-lg border border-zinc-800 bg-zinc-900/50">
      <table className="w-full text-left text-sm text-zinc-400">
        <thead className="bg-zinc-900 text-xs uppercase text-zinc-500">
          <tr>
            <th scope="col" className="px-6 py-4 font-medium">Status</th>
            <th scope="col" className="px-6 py-4 font-medium text-zinc-200">Title</th>
            <th scope="col" className="px-6 py-4 font-medium">Acceptance</th>
            <th scope="col" className="px-6 py-4 font-medium">Difficulty</th>
          </tr>
        </thead>
        <tbody className="divide-y divide-zinc-800">
          {problems.map((problem) => (
            <tr key={problem.id} className="transition-colors hover:bg-zinc-800/50">
              <td className="px-6 py-4">
                {problem.user_status === "solved" ? (
                  <span className="text-emerald-500">✓</span>
                ) : problem.user_status === "attempted" ? (
                  <span className="text-amber-500">~</span>
                ) : (
                  <span className="text-zinc-600">-</span>
                )}
              </td>
              <td className="px-6 py-4 font-medium text-zinc-200">
                <Link href={`/problems/${problem.slug}`} className="hover:text-indigo-400">
                  {problem.title}
                </Link>
                {problem.tags && problem.tags.length > 0 && (
                  <div className="mt-1 flex gap-2">
                    {problem.tags.map((t) => (
                      <span key={t.id} className="text-[10px] text-zinc-500">{t.name}</span>
                    ))}
                  </div>
                )}
              </td>
              <td className="px-6 py-4">{problem.acceptance_rate.toFixed(1)}%</td>
              <td className="px-6 py-4">
                <DifficultyBadge difficulty={problem.difficulty} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
