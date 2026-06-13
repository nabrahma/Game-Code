"use client";

import React from "react";
import { Panel, PanelGroup, PanelResizeHandle } from "react-resizable-panels";

export default function SolveLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="h-[calc(100vh-64px)] w-full overflow-hidden bg-[#0A0A0A] p-2">
      <PanelGroup direction="horizontal" className="h-full rounded-lg">
        
        {/* Left Panel: Description & Tabs */}
        <Panel defaultSize={40} minSize={25} className="flex flex-col overflow-hidden rounded-lg border border-zinc-800 bg-zinc-950">
          {children}
        </Panel>

        <PanelResizeHandle className="w-2 transition-colors hover:bg-indigo-500/50 flex flex-col justify-center items-center cursor-col-resize group">
            <div className="h-8 w-1 rounded bg-zinc-700 group-hover:bg-indigo-400" />
        </PanelResizeHandle>

        {/* Right Panels: Editor (Top) & Console (Bottom) - Placeholders for Phase 4 */}
        <Panel defaultSize={60} minSize={30}>
          <PanelGroup direction="vertical">
            <Panel defaultSize={70} minSize={20} className="rounded-lg border border-zinc-800 bg-zinc-950 p-4 flex flex-col justify-center items-center text-zinc-600">
                Phase 4: Code Editor
            </Panel>

            <PanelResizeHandle className="h-2 transition-colors hover:bg-indigo-500/50 flex justify-center items-center cursor-row-resize group">
                <div className="h-1 w-8 rounded bg-zinc-700 group-hover:bg-indigo-400" />
            </PanelResizeHandle>

            <Panel defaultSize={30} minSize={10} className="rounded-lg border border-zinc-800 bg-zinc-950 p-4 flex flex-col justify-center items-center text-zinc-600">
                Phase 4: Console Output & Run/Submit
            </Panel>
          </PanelGroup>
        </Panel>

      </PanelGroup>
    </div>
  );
}
