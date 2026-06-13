"use client";

import React, { useState } from "react";
import { useParams } from "next/navigation";
import { useProblem } from "@/lib/hooks/useProblems";
import { ProblemDetailHeader } from "@/components/problems/solve/ProblemDetailHeader";
import { ProblemDetailTabs } from "@/components/problems/solve/ProblemDetailTabs";
import { ProblemDescription } from "@/components/problems/solve/ProblemDescription";
import { ProblemExamples } from "@/components/problems/solve/ProblemExamples";

export default function ProblemSolvePage() {
  const params = useParams();
  const slug = params.slug as string;
  const { data: problem, isLoading, error } = useProblem(slug);

  const [activeTab, setActiveTab] = useState<"description" | "submissions" | "editorial">("description");

  if (isLoading) {
    return (
      <div className="flex h-full items-center justify-center text-zinc-500">
        <div className="h-6 w-6 animate-spin rounded-full border-2 border-indigo-500 border-t-transparent" />
      </div>
    );
  }

  if (error || !problem) {
    return (
      <div className="flex h-full items-center justify-center text-red-500">
        Failed to load problem data.
      </div>
    );
  }

  return (
    <div className="flex h-full flex-col bg-[#0A0A0A] overflow-hidden rounded-lg">
      <ProblemDetailHeader 
        title={problem.title} 
        slug={problem.slug} 
        difficulty={problem.difficulty}
        isFavorite={problem.is_favorite}
      />
      
      <ProblemDetailTabs activeTab={activeTab} setActiveTab={setActiveTab} />

      <div className="flex-1 overflow-y-auto p-6 scrollbar-thin scrollbar-track-transparent scrollbar-thumb-zinc-700">
        {activeTab === "description" && (
          <div className="pb-10">
            <ProblemDescription content={problem.description} />
            <ProblemExamples examples={problem.examples || []} />
            
            {problem.constraints && (
              <div className="mt-8">
                <h3 className="mb-3 text-sm font-semibold text-zinc-300">Constraints:</h3>
                <code className="block whitespace-pre-wrap rounded-md bg-zinc-900 p-3 text-sm text-zinc-400 border border-zinc-800">
                  {problem.constraints}
                </code>
              </div>
            )}
          </div>
        )}

        {activeTab === "submissions" && (
          <div className="flex h-full items-center justify-center text-zinc-500 text-sm">
            Please log in to view your submissions.
          </div>
        )}

        {activeTab === "editorial" && (
          <div className="flex h-full items-center justify-center text-zinc-500 text-sm">
            Editorial content will be available after solving or unlocking.
          </div>
        )}
      </div>
    </div>
  );
}
