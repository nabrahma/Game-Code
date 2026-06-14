"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import { ArrowLeft, Save, Code, Database, FileText, Settings } from "lucide-react";
import { TestCaseManager } from "./TestCaseManager";
import { StarterCodeEditor } from "./StarterCodeEditor";
import { EditorialEditor } from "./EditorialEditor";
import { useProblem } from "@/hooks/useProblem";

type Tab = "general" | "testcases" | "startercode" | "editorial";

export default function EditProblemPage({ params }: { params: { slug: string } }) {
  const { data: problem, isLoading, mutate } = useProblem(params.slug);
  const [activeTab, setActiveTab] = useState<Tab>("general");
  const [saving, setSaving] = useState(false);
  
  const [formData, setFormData] = useState({
    title: "",
    difficulty: "easy",
    description: "",
    constraints: "",
    status: "draft"
  });

  useEffect(() => {
    if (problem) {
      setFormData({
        title: problem.title || "",
        difficulty: problem.difficulty || "easy",
        description: problem.description || "",
        constraints: problem.constraints || "",
        status: problem.status || "draft"
      });
    }
  }, [problem]);

  const handleSaveGeneral = async (e: React.FormEvent) => {
    e.preventDefault();
    setSaving(true);
    try {
      const res = await fetch(`http://localhost:8080/api/admin/problems/${params.slug}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData)
      });
      if (res.ok) {
        mutate();
      }
    } catch (err) {
      console.error(err);
    } finally {
      setSaving(false);
    }
  };

  if (isLoading) return <div className="p-8 text-zinc-400">Loading problem data...</div>;
  if (!problem) return <div className="p-8 text-red-400">Problem not found</div>;

  return (
    <div className="flex flex-col h-full bg-[#0A0A0A]">
      <div className="border-b border-zinc-800 bg-zinc-950 p-6 sticky top-0 z-10">
        <div className="max-w-6xl mx-auto flex items-center justify-between">
          <div>
            <Link 
              href="/admin/problems"
              className="inline-flex items-center text-sm text-zinc-400 hover:text-white transition-colors mb-2"
            >
              <ArrowLeft className="w-4 h-4 mr-1" />
              Back to Problems
            </Link>
            <h1 className="text-2xl font-bold text-white flex items-center gap-3">
              {problem.title}
              <span className="text-xs font-mono bg-zinc-800 text-zinc-400 px-2 py-1 rounded-md">
                {problem.id}
              </span>
            </h1>
          </div>

          {activeTab === "general" && (
            <button
              onClick={handleSaveGeneral}
              disabled={saving}
              className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white px-4 py-2 rounded-lg font-medium transition-colors"
            >
              <Save className="w-4 h-4" />
              {saving ? "Saving..." : "Save Changes"}
            </button>
          )}
        </div>

        <div className="max-w-6xl mx-auto mt-6 flex gap-2">
          <TabButton active={activeTab === "general"} onClick={() => setActiveTab("general")} icon={<Settings className="w-4 h-4"/>}>General</TabButton>
          <TabButton active={activeTab === "testcases"} onClick={() => setActiveTab("testcases")} icon={<Database className="w-4 h-4"/>}>Test Cases</TabButton>
          <TabButton active={activeTab === "startercode"} onClick={() => setActiveTab("startercode")} icon={<Code className="w-4 h-4"/>}>Starter Code</TabButton>
          <TabButton active={activeTab === "editorial"} onClick={() => setActiveTab("editorial")} icon={<FileText className="w-4 h-4"/>}>Editorial</TabButton>
        </div>
      </div>

      <div className="flex-1 overflow-y-auto p-6">
        <div className="max-w-6xl mx-auto">
          {activeTab === "general" && (
            <form id="general-form" className="space-y-6">
              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <label className="block text-sm font-medium text-zinc-300">Title</label>
                  <input
                    type="text"
                    className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-indigo-500"
                    value={formData.title}
                    onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                  />
                </div>
                <div className="space-y-2">
                  <label className="block text-sm font-medium text-zinc-300">Slug (Read-only)</label>
                  <input
                    type="text"
                    className="w-full bg-zinc-950 border border-zinc-800 rounded-lg px-4 py-2.5 text-zinc-500 focus:outline-none"
                    value={problem.slug}
                    readOnly
                  />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <label className="block text-sm font-medium text-zinc-300">Difficulty</label>
                  <select
                    className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-indigo-500"
                    value={formData.difficulty}
                    onChange={(e) => setFormData({ ...formData, difficulty: e.target.value })}
                  >
                    <option value="easy">Easy</option>
                    <option value="medium">Medium</option>
                    <option value="hard">Hard</option>
                  </select>
                </div>
                
                <div className="space-y-2">
                  <label className="block text-sm font-medium text-zinc-300">Status</label>
                  <select
                    className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-indigo-500"
                    value={formData.status}
                    onChange={(e) => setFormData({ ...formData, status: e.target.value })}
                  >
                    <option value="draft">Draft</option>
                    <option value="review">Review</option>
                    <option value="published">Published</option>
                  </select>
                </div>
              </div>

              <div className="space-y-2">
                <label className="block text-sm font-medium text-zinc-300">Description (Markdown)</label>
                <textarea
                  className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-3 text-white focus:outline-none focus:border-indigo-500 h-96 font-mono text-sm resize-y"
                  value={formData.description}
                  onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                />
              </div>

              <div className="space-y-2">
                <label className="block text-sm font-medium text-zinc-300">Constraints</label>
                <textarea
                  className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-3 text-white focus:outline-none focus:border-indigo-500 h-24 font-mono text-sm resize-y"
                  value={formData.constraints}
                  onChange={(e) => setFormData({ ...formData, constraints: e.target.value })}
                />
              </div>
            </form>
          )}

          {activeTab === "testcases" && <TestCaseManager problemId={problem.id} initialTestCases={problem.test_cases || []} onUpdate={() => mutate()} />}
          {activeTab === "startercode" && <StarterCodeEditor problemId={problem.id} initialCodes={problem.starter_code || []} onUpdate={() => mutate()} />}
          {activeTab === "editorial" && <EditorialEditor problemId={problem.id} initialEditorial={problem.editorial} onUpdate={() => mutate()} />}
        </div>
      </div>
    </div>
  );
}

function TabButton({ active, onClick, children, icon }: { active: boolean, onClick: () => void, children: React.ReactNode, icon: React.ReactNode }) {
  return (
    <button
      onClick={onClick}
      className={`flex items-center gap-2 px-4 py-2 rounded-t-lg font-medium text-sm transition-colors border-b-2 ${
        active 
          ? "bg-zinc-900 border-indigo-500 text-indigo-400" 
          : "border-transparent text-zinc-400 hover:bg-zinc-900 hover:text-zinc-200"
      }`}
    >
      {icon}
      {children}
    </button>
  );
}
