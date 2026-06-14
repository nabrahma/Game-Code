'use client';

import React, { useEffect, useState } from 'react';
import Link from 'next/link';

interface ProblemList {
  id: string;
  title: string;
  slug: string;
  description: string;
  is_curated: boolean;
}

export default function ListsPage() {
  const [curatedLists, setCuratedLists] = useState<ProblemList[]>([]);
  const [userLists, setUserLists] = useState<ProblemList[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchLists() {
      try {
        const [curatedRes, userRes] = await Promise.all([
          fetch('http://localhost:8080/api/lists/curated'),
          fetch('http://localhost:8080/api/lists/user?user_id=00000000-0000-0000-0000-000000000001')
        ]);
        
        if (curatedRes.ok) {
          const data = await curatedRes.json();
          setCuratedLists(data);
        }
        if (userRes.ok) {
          const data = await userRes.json();
          setUserLists(data);
        }
      } catch (err) {
        console.error('Failed to fetch lists', err);
      } finally {
        setLoading(false);
      }
    }
    fetchLists();
  }, []);

  if (loading) {
    return <div className="p-8 text-white">Loading lists...</div>;
  }

  return (
    <div className="max-w-7xl mx-auto p-8 text-white space-y-12">
      <header>
        <h1 className="text-4xl font-bold bg-gradient-to-r from-blue-400 to-emerald-400 bg-clip-text text-transparent">
          Learning Paths & Lists
        </h1>
        <p className="text-zinc-400 mt-2 text-lg">Curated tracks to master game development algorithms.</p>
      </header>

      <section>
        <h2 className="text-2xl font-semibold mb-6 flex items-center">
          <span className="bg-emerald-500/20 text-emerald-400 p-2 rounded-lg mr-3">
            <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </span>
          Curated Paths
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {curatedLists.map((list) => (
            <Link key={list.id} href={`/lists/${list.slug}`}>
              <div className="bg-zinc-900 border border-zinc-800 p-6 rounded-xl hover:border-emerald-500/50 transition-all cursor-pointer group hover:-translate-y-1">
                <h3 className="text-xl font-bold text-white group-hover:text-emerald-400 transition-colors">{list.title}</h3>
                <p className="text-zinc-400 mt-3">{list.description}</p>
                <div className="mt-6 flex items-center text-sm font-medium text-emerald-500">
                  Start Path
                  <svg className="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                  </svg>
                </div>
              </div>
            </Link>
          ))}
          {curatedLists.length === 0 && (
            <div className="text-zinc-500 col-span-full">No curated lists available.</div>
          )}
        </div>
      </section>

      <section>
        <h2 className="text-2xl font-semibold mb-6 flex items-center">
          <span className="bg-blue-500/20 text-blue-400 p-2 rounded-lg mr-3">
            <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
            </svg>
          </span>
          Your Lists
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {userLists.map((list) => (
            <Link key={list.id} href={`/lists/${list.slug}`}>
              <div className="bg-zinc-900 border border-zinc-800 p-6 rounded-xl hover:border-blue-500/50 transition-all cursor-pointer group">
                <h3 className="text-xl font-bold text-white group-hover:text-blue-400 transition-colors">{list.title}</h3>
                <p className="text-zinc-400 mt-3">{list.description || 'No description'}</p>
              </div>
            </Link>
          ))}
          {userLists.length === 0 && (
            <div className="text-zinc-500 col-span-full border border-dashed border-zinc-800 rounded-xl p-8 text-center">
              You haven't created any lists yet. <br/> Go to a problem and click "Add to List" to start building your collection!
            </div>
          )}
        </div>
      </section>
    </div>
  );
}
