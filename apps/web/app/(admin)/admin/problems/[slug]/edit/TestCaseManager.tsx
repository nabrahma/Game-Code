import React, { useState } from "react";
import { Plus, Trash2, Save } from "lucide-react";

export function TestCaseManager({ problemId, initialTestCases, onUpdate }: { problemId: string, initialTestCases: any[], onUpdate: () => void }) {
  const [testCases, setTestCases] = useState(initialTestCases);
  const [saving, setSaving] = useState(false);

  const handleAdd = () => {
    setTestCases([...testCases, { id: "", problem_id: problemId, input: "", output: "", is_hidden: true, order_index: testCases.length, time_limit: 2000, memory_limit: 262144 }]);
  };

  const handleDelete = async (index: number, id: string) => {
    if (id) {
      if (!confirm("Delete this test case permanently?")) return;
      await fetch(`http://localhost:8080/api/admin/testcases/${id}`, { method: "DELETE" });
      onUpdate();
    }
    const newTc = [...testCases];
    newTc.splice(index, 1);
    setTestCases(newTc);
  };

  const handleSave = async (index: number) => {
    setSaving(true);
    try {
      const tc = testCases[index];
      await fetch(`http://localhost:8080/api/admin/testcases`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(tc)
      });
      onUpdate();
    } catch (err) {
      console.error(err);
    } finally {
      setSaving(false);
    }
  };

  const updateField = (index: number, field: string, value: any) => {
    const newTc = [...testCases];
    newTc[index] = { ...newTc[index], [field]: value };
    setTestCases(newTc);
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-xl font-bold text-white">Test Cases</h2>
          <p className="text-sm text-zinc-400">Manage input/output pairs for evaluating submissions.</p>
        </div>
        <button onClick={handleAdd} className="flex items-center gap-2 bg-zinc-800 hover:bg-zinc-700 text-white px-3 py-2 rounded-lg text-sm font-medium transition-colors">
          <Plus className="w-4 h-4" /> Add Test Case
        </button>
      </div>

      <div className="space-y-4">
        {testCases.map((tc, index) => (
          <div key={index} className="bg-zinc-900 border border-zinc-800 rounded-xl p-4">
            <div className="flex gap-4 mb-4">
              <div className="flex-1 space-y-2">
                <label className="text-xs font-medium text-zinc-400">Input</label>
                <textarea 
                  value={tc.input} 
                  onChange={(e) => updateField(index, "input", e.target.value)}
                  className="w-full bg-zinc-950 border border-zinc-800 rounded-lg p-3 text-white font-mono text-sm focus:border-indigo-500 focus:outline-none h-24"
                  placeholder="e.g. 1 2 3"
                />
              </div>
              <div className="flex-1 space-y-2">
                <label className="text-xs font-medium text-zinc-400">Expected Output</label>
                <textarea 
                  value={tc.output} 
                  onChange={(e) => updateField(index, "output", e.target.value)}
                  className="w-full bg-zinc-950 border border-zinc-800 rounded-lg p-3 text-white font-mono text-sm focus:border-indigo-500 focus:outline-none h-24"
                  placeholder="e.g. 6"
                />
              </div>
            </div>
            <div className="flex justify-between items-center bg-zinc-950/50 p-3 rounded-lg border border-zinc-800/50">
              <div className="flex gap-6 items-center">
                <label className="flex items-center gap-2 text-sm text-zinc-300">
                  <input type="checkbox" checked={tc.is_hidden} onChange={(e) => updateField(index, "is_hidden", e.target.checked)} className="rounded border-zinc-700 text-indigo-600 focus:ring-indigo-600 bg-zinc-900" />
                  Hidden from users
                </label>
                <label className="flex items-center gap-2 text-sm text-zinc-300">
                  Time Limit (ms):
                  <input type="number" value={tc.time_limit} onChange={(e) => updateField(index, "time_limit", parseInt(e.target.value))} className="bg-zinc-900 border border-zinc-700 rounded px-2 py-1 w-20 text-white" />
                </label>
              </div>
              <div className="flex gap-2">
                <button onClick={() => handleSave(index)} disabled={saving} className="flex items-center gap-2 text-indigo-400 hover:text-indigo-300 bg-indigo-500/10 hover:bg-indigo-500/20 px-3 py-1.5 rounded-lg text-sm font-medium transition-colors">
                  <Save className="w-4 h-4" /> Save
                </button>
                <button onClick={() => handleDelete(index, tc.id)} className="flex items-center gap-2 text-rose-400 hover:text-rose-300 bg-rose-500/10 hover:bg-rose-500/20 px-3 py-1.5 rounded-lg text-sm font-medium transition-colors">
                  <Trash2 className="w-4 h-4" /> Delete
                </button>
              </div>
            </div>
          </div>
        ))}
        {testCases.length === 0 && (
          <div className="text-center py-12 border border-dashed border-zinc-800 rounded-xl">
            <p className="text-zinc-500">No test cases found.</p>
          </div>
        )}
      </div>
    </div>
  );
}
