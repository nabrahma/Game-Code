import React, { useState } from "react";
import { Save } from "lucide-react";
import Editor from "@monaco-editor/react";

const SUPPORTED_LANGUAGES = [
  { id: "cpp", label: "C++", defaultCode: "class Solution {\npublic:\n    // Add methods here\n};" },
  { id: "csharp", label: "C#", defaultCode: "public class Solution {\n    // Add methods here\n}" },
  { id: "lua", label: "Lua", defaultCode: "-- Write your solution here" },
  { id: "gdscript", label: "GDScript", defaultCode: "class_name Solution\n\n# Write your solution here" },
];

export function StarterCodeEditor({ problemId, initialCodes, onUpdate }: { problemId: string, initialCodes: any[], onUpdate: () => void }) {
  const [activeLang, setActiveLang] = useState("cpp");
  const [codes, setCodes] = useState<Record<string, { id?: string; code: string }>>(() => {
    const map: Record<string, { id?: string; code: string }> = {};
    initialCodes.forEach(c => {
      map[c.language] = { id: c.id, code: c.code };
    });
    return map;
  });
  const [saving, setSaving] = useState(false);

  const handleEditorChange = (value: string | undefined) => {
    setCodes(prev => ({
      ...prev,
      [activeLang]: { ...prev[activeLang], code: value || "" }
    }));
  };

  const handleSave = async () => {
    setSaving(true);
    try {
      const codeData = codes[activeLang];
      const payload = {
        id: codeData?.id,
        problem_id: problemId,
        language: activeLang,
        code: codeData?.code || ""
      };

      await fetch(`http://localhost:8080/api/admin/startercode`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });
      onUpdate();
    } catch (err) {
      console.error(err);
    } finally {
      setSaving(false);
    }
  };

  const currentCode = codes[activeLang]?.code ?? SUPPORTED_LANGUAGES.find(l => l.id === activeLang)?.defaultCode ?? "";

  return (
    <div className="space-y-6 flex flex-col h-[600px]">
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-xl font-bold text-white">Starter Code</h2>
          <p className="text-sm text-zinc-400">Provide the boilerplate code users will see when they start.</p>
        </div>
        <button onClick={handleSave} disabled={saving} className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white px-4 py-2 rounded-lg font-medium transition-colors shadow-lg shadow-indigo-500/20">
          <Save className="w-4 h-4" /> {saving ? "Saving..." : "Save Code"}
        </button>
      </div>

      <div className="flex gap-2">
        {SUPPORTED_LANGUAGES.map(lang => (
          <button
            key={lang.id}
            onClick={() => setActiveLang(lang.id)}
            className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors ${
              activeLang === lang.id ? "bg-zinc-800 text-white" : "bg-zinc-900 text-zinc-400 hover:text-white"
            }`}
          >
            {lang.label}
          </button>
        ))}
      </div>

      <div className="flex-1 rounded-xl overflow-hidden border border-zinc-800">
        <Editor
          height="100%"
          language={activeLang === 'gdscript' ? 'python' : activeLang} // GDScript fallback for syntax
          theme="vs-dark"
          value={currentCode}
          onChange={handleEditorChange}
          options={{
            minimap: { enabled: false },
            fontSize: 14,
            padding: { top: 16 },
            scrollBeyondLastLine: false,
          }}
        />
      </div>
    </div>
  );
}
