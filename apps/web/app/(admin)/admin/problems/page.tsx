"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import { Plus, Edit2, Trash2 } from "lucide-react";
import { TableSkeleton } from "@/components/ui/Skeleton";
import { EmptyState } from "@/components/ui/EmptyState";
import { ErrorState } from "@/components/ui/ErrorState";

interface ProblemSummary {
  id: string;
  slug: string;
  title: string;
  difficulty: string;
  acceptance_rate: number;
}

export default function AdminProblemsPage() {
  const [problems, setProblems] = useState<ProblemSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchProblems();
  }, []);

  const fetchProblems = async () => {
    setLoading(true);
    setError(null);
    try {
      const res = await fetch("http://localhost:8080/api/problems");
      if (!res.ok) throw new Error("Failed to load problems");
      const data = await res.json();
      setProblems(data.items || []);
    } catch (err: any) {
      console.error("Failed to fetch problems", err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id: string) => {
    if (!confirm("Are you sure you want to delete this problem?")) return;
    try {
      const res = await fetch(`http://localhost:8080/api/admin/problems/${id}`, {
        method: 'DELETE'
      });
      if (res.ok) {
        fetchProblems();
      }
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="p-8">
      <div className="flex items-center justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-white tracking-tight">Problems Database</h1>
          <p className="text-zinc-400 mt-2">Manage all coding challenges, test cases, and starter templates.</p>
        </div>
        <Link 
          href="/admin/problems/new"
          className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg font-medium transition-colors shadow-lg shadow-indigo-500/20"
        >
          <Plus className="w-5 h-5" />
          Create Problem
        </Link>
      </div>

      {loading ? (
        <TableSkeleton />
      ) : error ? (
        <ErrorState message={error} onRetry={fetchProblems} />
      ) : problems.length === 0 ? (
        <EmptyState 
          title="No problems found" 
          message="Click 'Create Problem' to add your first coding challenge." 
          action={
            <Link 
              href="/admin/problems/new"
              className="inline-flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg font-medium transition-colors"
            >
              <Plus className="w-4 h-4" /> Create Problem
            </Link>
          }
        />
      ) : (
        <div className="bg-zinc-900 border border-zinc-800 rounded-xl overflow-hidden shadow-2xl">
          <table className="w-full text-left text-sm text-zinc-400">
            <thead className="bg-zinc-950 border-b border-zinc-800 text-zinc-300 uppercase font-medium text-xs">
              <tr>
                <th className="px-6 py-4">Title</th>
                <th className="px-6 py-4">Difficulty</th>
                <th className="px-6 py-4">Acceptance</th>
                <th className="px-6 py-4 text-right">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-zinc-800/50">
              {problems.map((p) => (
                <tr key={p.id} className="hover:bg-zinc-800/30 transition-colors">
                  <td className="px-6 py-4 font-medium text-zinc-200">
                    {p.title}
                    <div className="text-xs text-zinc-500 font-mono mt-1">{p.slug}</div>
                  </td>
                  <td className="px-6 py-4">
                    <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium border ${
                      p.difficulty === 'easy' ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' :
                      p.difficulty === 'medium' ? 'bg-yellow-500/10 text-yellow-400 border-yellow-500/20' :
                      'bg-rose-500/10 text-rose-400 border-rose-500/20'
                    }`}>
                      {p.difficulty}
                    </span>
                  </td>
                  <td className="px-6 py-4 font-mono text-zinc-400">
                    {p.acceptance_rate.toFixed(1)}%
                  </td>
                  <td className="px-6 py-4 text-right space-x-2">
                    <Link 
                      href={`/admin/problems/${p.slug}/edit`}
                      className="inline-flex p-2 text-zinc-400 hover:text-indigo-400 hover:bg-indigo-400/10 rounded transition-colors"
                    >
                      <Edit2 className="w-4 h-4" />
                    </Link>
                    <button 
                      onClick={() => handleDelete(p.id)}
                      className="inline-flex p-2 text-zinc-400 hover:text-rose-400 hover:bg-rose-400/10 rounded transition-colors"
                    >
                      <Trash2 className="w-4 h-4" />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
