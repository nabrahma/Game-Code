"use client";

import React, { useState, useEffect } from "react";
import { useParams } from "next/navigation";
import { useProblem } from "@/lib/hooks/useProblems";
import { useExecution } from "@/lib/hooks/useExecution";
import { useSubmission } from "@/lib/hooks/useSubmission";
import { ProblemDetailHeader } from "@/components/problems/solve/ProblemDetailHeader";
import { ProblemDetailTabs } from "@/components/problems/solve/ProblemDetailTabs";
import { ProblemDescription } from "@/components/problems/solve/ProblemDescription";
import { ProblemExamples } from "@/components/problems/solve/ProblemExamples";
import { EditorHeader } from "@/components/problems/solve/EditorHeader";
import { MonacoWrapper } from "@/components/problems/solve/MonacoWrapper";
import { ConsolePanel } from "@/components/problems/solve/ConsolePanel";
import { DiscussFeed } from "@/components/problems/solve/DiscussFeed";
import { Panel, PanelGroup, PanelResizeHandle } from "react-resizable-panels";

export default function ProblemSolvePage() {
  const params = useParams();
  const slug = params.slug as string;
  const { data: problem, isLoading, error } = useProblem(slug);

  const [activeTab, setActiveTab] = useState<"description" | "submissions" | "editorial" | "discuss">("description");
  const [language, setLanguage] = useState("cpp");
  const [code, setCode] = useState("");
  const [input, setInput] = useState("");

  const { executeCode, isExecuting, output, status, setOutput } = useExecution();
  const { submit, isSubmitting, submission } = useSubmission();

  // Initialize starter code when problem loads or language changes
  useEffect(() => {
    if (problem?.starter_code) {
      const starter = problem.starter_code.find((s: any) => s.language === language);
      if (starter) {
        // In a real app we'd check localStorage first
        setCode(starter.code);
      }
    }
  }, [problem, language]);

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

  const handleRun = () => {
    if (!code) return;
    executeCode(problem.id, language, code, input);
  };

  const handleSubmit = () => {
    if (!code) return;
    submit({ problem_id: problem.id, language, code });
  };

  return (
    <PanelGroup direction="horizontal" className="h-full rounded-lg">
      {/* Left Panel: Description */}
      <Panel defaultSize={40} minSize={25} className="flex flex-col overflow-hidden rounded-lg border border-zinc-800 bg-[#0A0A0A]">
        <ProblemDetailHeader 
          problemId={problem.id}
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
          {activeTab === "discuss" && (
            <DiscussFeed slug={slug} />
          )}
        </div>
      </Panel>

      <PanelResizeHandle className="w-2 transition-colors hover:bg-indigo-500/50 flex flex-col justify-center items-center cursor-col-resize group">
        <div className="h-8 w-1 rounded bg-zinc-700 group-hover:bg-indigo-400" />
      </PanelResizeHandle>

      {/* Right Panels: Editor & Console */}
      <Panel defaultSize={60} minSize={30}>
        <PanelGroup direction="vertical">
          <Panel defaultSize={65} minSize={20} className="flex flex-col overflow-hidden rounded-lg border border-zinc-800 bg-zinc-950">
            <EditorHeader 
              language={language}
              setLanguage={setLanguage}
              onRun={handleRun}
              onSubmit={handleSubmit}
              isExecuting={isExecuting || isSubmitting}
            />
            <div className="flex-1">
              <MonacoWrapper 
                language={language} 
                code={code} 
                onChange={(val) => setCode(val || "")} 
              />
            </div>
          </Panel>

          <PanelResizeHandle className="h-2 transition-colors hover:bg-indigo-500/50 flex justify-center items-center cursor-row-resize group">
            <div className="h-1 w-8 rounded bg-zinc-700 group-hover:bg-indigo-400" />
          </PanelResizeHandle>

          <Panel defaultSize={35} minSize={10} className="overflow-hidden rounded-lg border border-zinc-800 bg-zinc-950">
            <ConsolePanel 
              status={isSubmitting ? (submission?.verdict || "submitting") : status}
              output={output}
              input={input}
              setInput={setInput}
              submission={submission}
            />
          </Panel>
        </PanelGroup>
      </Panel>
    </PanelGroup>
  );
}
