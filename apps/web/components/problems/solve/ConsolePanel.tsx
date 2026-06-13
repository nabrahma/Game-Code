import React from "react";
import { Terminal, Loader2 } from "lucide-react";
import clsx from "clsx";

interface ConsolePanelProps {
  status: string;
  output: string;
  input: string;
  setInput: (val: string) => void;
}

export function ConsolePanel({ status, output, input, setInput }: ConsolePanelProps) {
  return (
    <div className="flex h-full flex-col bg-zinc-950">
      <div className="flex items-center gap-2 border-b border-zinc-800 bg-zinc-900 px-4 py-2 text-sm font-medium text-zinc-300">
        <Terminal className="h-4 w-4" />
        Console
        {status && (
          <span className="ml-auto flex items-center gap-2 text-xs">
            {["queued", "submitting", "connecting...", "running"].includes(status) && (
              <Loader2 className="h-3 w-3 animate-spin text-indigo-400" />
            )}
            <span
              className={clsx({
                "text-emerald-400": status === "success",
                "text-red-400": status === "error" || status === "timeout",
                "text-indigo-400": !["success", "error", "timeout"].includes(status),
              })}
            >
              {status.toUpperCase()}
            </span>
          </span>
        )}
      </div>

      <div className="flex flex-1 gap-4 overflow-hidden p-4">
        <div className="flex w-1/2 flex-col">
          <label className="mb-2 text-xs font-semibold uppercase tracking-wider text-zinc-500">
            Custom Testcase
          </label>
          <textarea
            value={input}
            onChange={(e) => setInput(e.target.value)}
            className="flex-1 resize-none rounded-md border border-zinc-800 bg-zinc-900/50 p-3 text-sm text-zinc-300 font-mono focus:border-indigo-500 focus:outline-none"
            placeholder="Enter custom input here..."
            spellCheck={false}
          />
        </div>

        <div className="flex w-1/2 flex-col">
          <label className="mb-2 text-xs font-semibold uppercase tracking-wider text-zinc-500">
            Output
          </label>
          <div className="flex-1 overflow-auto rounded-md border border-zinc-800 bg-zinc-900/50 p-3 text-sm text-zinc-300 font-mono">
            {output ? (
              <pre className="whitespace-pre-wrap">{output}</pre>
            ) : (
              <div className="text-zinc-600">Run code to see output</div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
