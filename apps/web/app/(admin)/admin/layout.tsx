import React from "react";
import Link from "next/link";
import { Settings, Database, Users } from "lucide-react";

export default function AdminLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-screen bg-[#0A0A0A] text-white">
      {/* Sidebar */}
      <aside className="w-64 border-r border-zinc-800 bg-zinc-950 flex flex-col">
        <div className="p-6">
          <Link href="/admin/problems" className="flex items-center gap-2">
            <div className="h-8 w-8 rounded bg-indigo-600 flex items-center justify-center font-bold text-white shadow-lg shadow-indigo-500/20">
              GC
            </div>
            <span className="text-xl font-bold tracking-tight">Admin CMS</span>
          </Link>
        </div>

        <nav className="flex-1 px-4 space-y-2">
          <Link 
            href="/admin/problems" 
            className="flex items-center gap-3 px-3 py-2 rounded-lg bg-zinc-900 text-indigo-400 font-medium transition-colors"
          >
            <Database className="w-5 h-5" />
            Problems
          </Link>
          <button 
            className="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-zinc-400 hover:bg-zinc-900 hover:text-white transition-colors cursor-not-allowed opacity-50"
          >
            <Users className="w-5 h-5" />
            Users
          </button>
          <button 
            className="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-zinc-400 hover:bg-zinc-900 hover:text-white transition-colors cursor-not-allowed opacity-50"
          >
            <Settings className="w-5 h-5" />
            Settings
          </button>
        </nav>
      </aside>

      {/* Main Content */}
      <main className="flex-1 overflow-y-auto bg-[#0A0A0A]">
        {children}
      </main>
    </div>
  );
}
