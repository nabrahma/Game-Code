"use client";

import { Suspense, useState, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { useProblems } from "@/lib/hooks/useProblems";
import { ProblemFiltersBar } from "@/components/problems/ProblemFiltersBar";
import { ProblemTable } from "@/components/problems/ProblemTable";
import { ErrorState } from "@/components/ui/ErrorState";

function ProblemListContent() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const urlQuery = searchParams.get("q") || "";
  const urlDifficulty = searchParams.get("difficulty") || "";

  const [searchQuery, setSearchQuery] = useState(urlQuery);
  const [difficulty, setDifficulty] = useState(urlDifficulty);

  // Debounce search query update to URL
  useEffect(() => {
    const timeoutId = setTimeout(() => {
      const params = new URLSearchParams(searchParams.toString());
      if (searchQuery) {
        params.set("q", searchQuery);
      } else {
        params.delete("q");
      }
      
      if (difficulty) {
        params.set("difficulty", difficulty);
      } else {
        params.delete("difficulty");
      }

      router.push(`/problems?${params.toString()}`, { scroll: false });
    }, 300);

    return () => clearTimeout(timeoutId);
  }, [searchQuery, difficulty, router, searchParams]);

  const { data, isLoading, error } = useProblems({
    q: searchQuery,
    difficulty: difficulty,
    limit: 50,
  });

  return (
    <div className="mx-auto max-w-6xl px-4 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-white">Problem Set</h1>
        <p className="mt-2 text-zinc-400">
          Master game engine architecture through structured practice.
        </p>
      </div>

      <ProblemFiltersBar
        searchQuery={searchQuery}
        setSearchQuery={setSearchQuery}
        difficulty={difficulty}
        setDifficulty={setDifficulty}
      />

      {error ? (
        <ErrorState message="Failed to load problems. Please try again later." />
      ) : (
        <ProblemTable problems={data?.problems || []} isLoading={isLoading} />
      )}
    </div>
  );
}

export default function ProblemsPage() {
  return (
    <Suspense fallback={<div className="p-8 text-center text-zinc-500">Loading...</div>}>
      <ProblemListContent />
    </Suspense>
  );
}
