import React from "react";
import clsx from "clsx";
import { FileText, History, Lightbulb, MessageCircle } from "lucide-react";

type Tab = "description" | "submissions" | "editorial" | "discuss";

interface ProblemDetailTabsProps {
  activeTab: Tab;
  setActiveTab: (tab: Tab) => void;
}

export function ProblemDetailTabs({ activeTab, setActiveTab }: ProblemDetailTabsProps) {
  const tabs = [
    { id: "description" as Tab, label: "Description", icon: FileText },
    { id: "submissions" as Tab, label: "Submissions", icon: History },
    { id: "discuss" as Tab, label: "Discuss", icon: MessageCircle },
    { id: "editorial" as Tab, label: "Editorial", icon: Lightbulb },
  ];

  return (
    <div className="flex border-b border-zinc-800 bg-zinc-900/50 px-2 pt-2">
      {tabs.map((tab) => {
        const Icon = tab.icon;
        const isActive = activeTab === tab.id;
        return (
          <button
            key={tab.id}
            onClick={() => setActiveTab(tab.id)}
            className={clsx(
              "flex items-center gap-2 rounded-t-lg px-4 py-2.5 text-sm font-medium transition-colors",
              isActive
                ? "bg-zinc-800 text-white"
                : "text-zinc-400 hover:bg-zinc-800/50 hover:text-zinc-200"
            )}
          >
            <Icon className="h-4 w-4" />
            {tab.label}
          </button>
        );
      })}
    </div>
  );
}
