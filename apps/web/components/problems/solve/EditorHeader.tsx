import React from "react";
import { Play, Send } from "lucide-react";

interface EditorHeaderProps {
  language: string;
  setLanguage: (lang: string) => void;
  onRun: () => void;
  onSubmit: () => void;
  isExecuting: boolean;
}

export function EditorHeader({
  language,
  setLanguage,
  onRun,
  onSubmit,
  isExecuting,
}: EditorHeaderProps) {
  return (
    <div className="flex h-12 shrink-0 items-center justify-between border-b border-zinc-800 bg-zinc-900 px-4">
      <div className="flex items-center gap-2">
        <select
          value={language}
          onChange={(e) => setLanguage(e.target.value)}
          className="rounded-md border border-zinc-800 bg-zinc-950 py-1 pl-2 pr-8 text-sm text-zinc-300 focus:border-indigo-500 focus:outline-none"
        >
          <option value="cpp">C++</option>
          <option value="csharp">C#</option>
          <option value="gdscript">GDScript</option>
          <option value="lua">Lua</option>
        </select>
      </div>

      <div className="flex items-center gap-2">
        <button
          onClick={onRun}
          disabled={isExecuting}
          className="flex items-center gap-2 rounded-md bg-zinc-800 px-3 py-1.5 text-sm font-semibold text-zinc-300 transition-colors hover:bg-zinc-700 disabled:opacity-50"
        >
          <Play className="h-4 w-4 text-emerald-400" />
          Run
        </button>
        <button
          onClick={onSubmit}
          disabled={isExecuting}
          className="flex items-center gap-2 rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold text-white transition-colors hover:bg-indigo-500 disabled:opacity-50"
        >
          <Send className="h-4 w-4" />
          Submit
        </button>
      </div>
    </div>
  );
}
