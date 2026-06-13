import React from "react";

type Difficulty = "easy" | "medium" | "hard";

export function DifficultyBadge({ difficulty }: { difficulty: Difficulty }) {
  const styles = {
    easy: "bg-emerald-500/10 text-emerald-500 border-emerald-500/20",
    medium: "bg-amber-500/10 text-amber-500 border-amber-500/20",
    hard: "bg-red-500/10 text-red-500 border-red-500/20",
  };

  return (
    <span
      className={`inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold ${styles[difficulty]}`}
    >
      {difficulty.charAt(0).toUpperCase() + difficulty.slice(1)}
    </span>
  );
}
