import React, { useState } from "react";
import { Save } from "lucide-react";

export function EditorialEditor({ problemId, initialEditorial, onUpdate }: { problemId: string, initialEditorial: any, onUpdate: () => void }) {
  const [saving, setSaving] = useState(false);
  const [formData, setFormData] = useState({
    id: initialEditorial?.id || undefined,
    problem_id: problemId,
    content: initialEditorial?.content || "",
    time_complexity: initialEditorial?.time_complexity || "",
    space_complexity: initialEditorial?.space_complexity || "",
    unity_variant: initialEditorial?.unity_variant || "",
    unreal_variant: initialEditorial?.unreal_variant || "",
    godot_variant: initialEditorial?.godot_variant || ""
  });

  const handleSave = async (e: React.FormEvent) => {
    e.preventDefault();
    setSaving(true);
    try {
      await fetch(`http://localhost:8080/api/admin/editorials`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData)
      });
      onUpdate();
    } catch (err) {
      console.error(err);
    } finally {
      setSaving(false);
    }
  };

  return (
    <form onSubmit={handleSave} className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-xl font-bold text-white">Editorial & Official Solution</h2>
          <p className="text-sm text-zinc-400">Write the detailed explanation for solving this problem.</p>
        </div>
        <button type="submit" disabled={saving} className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white px-4 py-2 rounded-lg font-medium transition-colors shadow-lg shadow-indigo-500/20">
          <Save className="w-4 h-4" /> {saving ? "Saving..." : "Save Editorial"}
        </button>
      </div>

      <div className="grid grid-cols-2 gap-6">
        <div className="space-y-2">
          <label className="block text-sm font-medium text-zinc-300">Time Complexity</label>
          <input
            type="text"
            className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2 text-white focus:outline-none focus:border-indigo-500 font-mono text-sm"
            placeholder="e.g. O(N log N)"
            value={formData.time_complexity}
            onChange={(e) => setFormData({ ...formData, time_complexity: e.target.value })}
          />
        </div>
        <div className="space-y-2">
          <label className="block text-sm font-medium text-zinc-300">Space Complexity</label>
          <input
            type="text"
            className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2 text-white focus:outline-none focus:border-indigo-500 font-mono text-sm"
            placeholder="e.g. O(N)"
            value={formData.space_complexity}
            onChange={(e) => setFormData({ ...formData, space_complexity: e.target.value })}
          />
        </div>
      </div>

      <div className="space-y-2">
        <label className="block text-sm font-medium text-zinc-300">Main Content (Markdown)</label>
        <textarea
          className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-3 text-white focus:outline-none focus:border-indigo-500 h-96 font-mono text-sm resize-y"
          placeholder="Explain the approach..."
          value={formData.content}
          onChange={(e) => setFormData({ ...formData, content: e.target.value })}
        />
      </div>

      <div className="space-y-4 pt-4 border-t border-zinc-800">
        <h3 className="text-lg font-medium text-white">Engine-Specific Notes (Optional)</h3>
        
        <div className="space-y-2">
          <label className="block text-sm text-zinc-400">Unity specific gotchas</label>
          <input type="text" className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2 text-white focus:border-indigo-500" value={formData.unity_variant} onChange={(e) => setFormData({ ...formData, unity_variant: e.target.value })} />
        </div>
        <div className="space-y-2">
          <label className="block text-sm text-zinc-400">Unreal specific gotchas</label>
          <input type="text" className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2 text-white focus:border-indigo-500" value={formData.unreal_variant} onChange={(e) => setFormData({ ...formData, unreal_variant: e.target.value })} />
        </div>
        <div className="space-y-2">
          <label className="block text-sm text-zinc-400">Godot specific gotchas</label>
          <input type="text" className="w-full bg-zinc-900 border border-zinc-800 rounded-lg px-4 py-2 text-white focus:border-indigo-500" value={formData.godot_variant} onChange={(e) => setFormData({ ...formData, godot_variant: e.target.value })} />
        </div>
      </div>
    </form>
  );
}
