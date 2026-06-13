import React from "react";
import { Terminal, Loader2 } from "lucide-react";
import clsx from "clsx";

import { Submission } from "@/lib/hooks/useSubmission";

interface ConsolePanelProps {
  status: string;
  output: string;
  input: string;
  setInput: (val: string) => void;
  submission?: Submission | null;
}

export function ConsolePanel({ status, output, input, setInput, submission }: ConsolePanelProps) {
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
                "text-emerald-400": status === "success" || submission?.verdict === "accepted",
                "text-red-400": status === "error" || status === "timeout" || (submission && submission.verdict !== "accepted" && submission.verdict !== "pending"),
                "text-indigo-400": !["success", "error", "timeout"].includes(status) && (!submission || submission.verdict === "pending"),
              })}
            >
              {submission ? `SUBMISSION: ${submission.verdict.toUpperCase()}` : status.toUpperCase()}
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
            {submission ? (
              <div className="space-y-2">
                <div className="font-bold text-lg mb-4 text-white">Verdict: {submission.verdict.toUpperCase()}</div>
                {submission.error_message && <div className="text-red-400">{submission.error_message}</div>}
                {submission.passed_test_count !== undefined && (
                  <div className="text-zinc-400">Tests passed: {submission.passed_test_count} / {submission.total_test_count}</div>
                )}
                {submission.results?.map((res, i) => (
                  <div key={i} className="mt-4 p-2 border border-zinc-800 rounded bg-zinc-950">
                    <div className="font-semibold text-xs text-zinc-500 mb-1">TEST CASE {i+1}</div>
                    <div className="mb-1">Passed: {res.passed ? "✅" : "❌"}</div>
                    <div className="mb-1">Input: <span className="text-zinc-400">{res.input}</span></div>
                    <div className="mb-1">Expected: <span className="text-emerald-400">{res.expected}</span></div>
                    <div className="mb-1">Actual: <span className={res.passed ? "text-emerald-400" : "text-red-400"}>{res.actual}</span></div>
                  </div>
                ))}
              </div>
            ) : output ? (
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
