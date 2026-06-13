import React from "react";

interface Example {
  id: string;
  input: string;
  output: string;
  explanation?: string;
}

interface ProblemExamplesProps {
  examples: Example[];
}

export function ProblemExamples({ examples }: ProblemExamplesProps) {
  if (!examples || examples.length === 0) return null;

  return (
    <div className="mt-8 space-y-6">
      {examples.map((ex, idx) => (
        <div key={ex.id} className="rounded-lg border border-zinc-800 bg-zinc-900/30 p-4">
          <p className="mb-3 font-semibold text-zinc-200">Example {idx + 1}:</p>
          <div className="space-y-2 text-sm">
            <div>
              <span className="font-semibold text-zinc-400">Input:</span>{" "}
              <code className="text-zinc-200">{ex.input}</code>
            </div>
            <div>
              <span className="font-semibold text-zinc-400">Output:</span>{" "}
              <code className="text-zinc-200">{ex.output}</code>
            </div>
            {ex.explanation && (
              <div className="mt-2 text-zinc-400">
                <span className="font-semibold">Explanation:</span> {ex.explanation}
              </div>
            )}
          </div>
        </div>
      ))}
    </div>
  );
}
