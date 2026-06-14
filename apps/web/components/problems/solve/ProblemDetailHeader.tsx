import React from "react";
import Link from "next/link";
import { ChevronLeft, ChevronRight, LayoutList } from "lucide-react";
import { DifficultyBadge } from "@/components/ui/DifficultyBadge";
import { FavoriteToggle } from "./FavoriteToggle";
import { AddToListDropdown } from "./AddToListDropdown";

interface ProblemDetailHeaderProps {
  problemId: string;
  title: string;
  slug: string;
  difficulty: "easy" | "medium" | "hard";
  isFavorite?: boolean;
}

export function ProblemDetailHeader({ problemId, title, slug, difficulty, isFavorite }: ProblemDetailHeaderProps) {
  return (
    <div className="flex items-center justify-between border-b border-zinc-800 bg-zinc-900 px-4 py-3">
      <div className="flex items-center gap-4">
        <Link 
          href="/problems"
          className="flex h-8 w-8 items-center justify-center rounded-md text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-white"
          title="Problem List"
        >
          <LayoutList className="h-5 w-5" />
        </Link>
        <h1 className="text-lg font-semibold text-zinc-100">{title}</h1>
        <DifficultyBadge difficulty={difficulty} />
      </div>

      <div className="flex items-center gap-2">
        <AddToListDropdown problemId={problemId} />
        <FavoriteToggle slug={slug} initialIsFavorite={isFavorite} />
        
        <div className="flex items-center gap-1 border-l border-zinc-800 pl-4">
          <button className="flex h-8 w-8 items-center justify-center rounded-md text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-white" title="Previous Problem">
            <ChevronLeft className="h-5 w-5" />
          </button>
          <button className="flex h-8 w-8 items-center justify-center rounded-md text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-white" title="Next Problem">
            <ChevronRight className="h-5 w-5" />
          </button>
        </div>
      </div>
    </div>
  );
}
