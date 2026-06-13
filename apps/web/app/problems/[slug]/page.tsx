"use client";

import { useProblem } from "@/lib/hooks/useProblems";
import { useParams } from "next/navigation";
import { DifficultyBadge } from "@/components/ui/DifficultyBadge";

export default function ProblemDetailPage() {
  const params = useParams();
  const slug = params.slug as string;
  
  const { data: problem, isLoading, error } = useProblem(slug);

  if (isLoading) {
    return <div className="p-8 text-center text-zinc-500">Loading problem details...</div>;
  }

  if (error || !problem) {
    return <div className="p-8 text-center text-red-500">Failed to load problem. It might not exist.</div>;
  }

  return (
    <div className="mx-auto max-w-4xl px-4 py-8">
      <div className="mb-6 flex items-center justify-between">
        <h1 className="text-3xl font-bold text-white">{problem.title}</h1>
        <DifficultyBadge difficulty={problem.difficulty} />
      </div>

      <div className="prose prose-invert max-w-none rounded-lg border border-zinc-800 bg-zinc-900/50 p-6">
        <p className="text-zinc-300">{problem.description}</p>
        
        <h3 className="mt-8 text-lg font-semibold text-white">Constraints</h3>
        <code className="block rounded bg-zinc-950 p-3 text-sm text-zinc-400">
          {problem.constraints}
        </code>
      </div>

      <div className="mt-8 rounded-lg border border-zinc-800 bg-zinc-900/50 p-8 text-center">
        <p className="text-zinc-500">
          Phase 3: Code Editor and Execution Environment will go here.
        </p>
      </div>
    </div>
  );
}
