'use client';

import React, { useEffect, useState } from 'react';
import Link from 'next/link';

interface ProblemSummary {
  id: string;
  slug: string;
  title: string;
  difficulty: string;
  acceptance_rate: number;
}

interface ProblemListItem {
  list_id: string;
  problem_id: string;
  order_index: number;
  problem: ProblemSummary;
}

interface ProblemList {
  id: string;
  title: string;
  slug: string;
  description: string;
  is_curated: boolean;
  items: ProblemListItem[];
}

export default function ListDetailPage({ params }: { params: { slug: string } }) {
  const [list, setList] = useState<ProblemList | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    async function fetchList() {
      try {
        const res = await fetch(`http://localhost:8080/api/lists/${params.slug}`);
        if (!res.ok) {
          throw new Error('List not found');
        }
        const data = await res.json();
        setList(data);
      } catch (err) {
        setError('Failed to load list. It might not exist.');
      } finally {
        setLoading(false);
      }
    }
    fetchList();
  }, [params.slug]);

  if (loading) return <div className="p-8 text-white">Loading list details...</div>;
  if (error || !list) return <div className="p-8 text-red-400">{error}</div>;

  return (
    <div className="max-w-5xl mx-auto p-8 text-white space-y-8">
      <header className="mb-10">
        <div className="flex items-center space-x-3 text-emerald-400 mb-4">
          <Link href="/lists" className="hover:text-emerald-300 transition-colors flex items-center">
            <svg className="w-4 h-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
            Back to Lists
          </Link>
        </div>
        <h1 className="text-4xl font-bold bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
          {list.title}
        </h1>
        <p className="text-zinc-400 mt-4 text-lg">{list.description}</p>
        <div className="mt-4 flex space-x-4 text-sm font-medium">
          <span className="bg-zinc-800 text-zinc-300 px-3 py-1 rounded-full">
            {list.items?.length || 0} Problems
          </span>
          {list.is_curated && (
            <span className="bg-emerald-500/20 text-emerald-400 px-3 py-1 rounded-full flex items-center">
              <svg className="w-4 h-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              Curated by GameCode
            </span>
          )}
        </div>
      </header>

      <div className="bg-zinc-900 border border-zinc-800 rounded-xl overflow-hidden">
        <table className="w-full text-left border-collapse">
          <thead>
            <tr className="bg-zinc-950 border-b border-zinc-800">
              <th className="p-4 font-semibold text-zinc-400 w-16 text-center">#</th>
              <th className="p-4 font-semibold text-zinc-400">Problem</th>
              <th className="p-4 font-semibold text-zinc-400 w-32 text-center">Difficulty</th>
              <th className="p-4 font-semibold text-zinc-400 w-32 text-right">Acceptance</th>
            </tr>
          </thead>
          <tbody>
            {!list.items || list.items.length === 0 ? (
              <tr>
                <td colSpan={4} className="p-8 text-center text-zinc-500">
                  This list is empty.
                </td>
              </tr>
            ) : (
              list.items.map((item, idx) => (
                <tr key={item.problem_id} className="border-b border-zinc-800/50 hover:bg-zinc-800/30 transition-colors">
                  <td className="p-4 text-center text-zinc-500 font-medium">{idx + 1}</td>
                  <td className="p-4">
                    <Link href={`/problems/${item.problem.slug}`} className="text-emerald-400 font-medium hover:underline">
                      {item.problem.title}
                    </Link>
                  </td>
                  <td className="p-4 text-center">
                    <span className={`px-2 py-1 text-xs font-bold uppercase rounded-md ${
                      item.problem.difficulty === 'easy' ? 'bg-green-500/10 text-green-400' :
                      item.problem.difficulty === 'medium' ? 'bg-yellow-500/10 text-yellow-400' :
                      'bg-red-500/10 text-red-400'
                    }`}>
                      {item.problem.difficulty}
                    </span>
                  </td>
                  <td className="p-4 text-right text-zinc-400">
                    {item.problem.acceptance_rate}%
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
