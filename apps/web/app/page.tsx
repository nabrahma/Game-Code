import Link from "next/link";

export default function Home() {
  return (
    <div className="min-h-screen flex flex-col bg-zinc-950 font-sans selection:bg-indigo-500/30">
      {/* Navbar */}
      <nav className="w-full border-b border-white/5 bg-zinc-950/50 backdrop-blur-md sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
          <div className="flex items-center gap-2">
            <div className="w-8 h-8 rounded bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center font-bold text-white shadow-lg shadow-indigo-500/20">
              G
            </div>
            <span className="text-xl font-bold tracking-tight text-white">GameCode</span>
          </div>
          <div className="flex items-center gap-6 text-sm font-medium text-zinc-400">
            <Link href="/problems" className="hover:text-white transition-colors">Problems</Link>
            <Link href="/leaderboard" className="hover:text-white transition-colors">Leaderboard</Link>
            <Link href="/login" className="px-4 py-2 rounded-md bg-white text-black hover:bg-zinc-200 transition-colors font-semibold">
              Sign In
            </Link>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <main className="flex-1 flex flex-col items-center justify-center relative overflow-hidden">
        {/* Background glow */}
        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-indigo-600/20 rounded-full blur-[120px] pointer-events-none" />
        
        <div className="z-10 text-center px-6 max-w-4xl mx-auto flex flex-col items-center">
          <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full border border-white/10 bg-white/5 text-xs font-medium text-indigo-300 mb-8 backdrop-blur-sm">
            <span className="w-2 h-2 rounded-full bg-indigo-400 animate-pulse" />
            Go Backend & Judge0 Migration Complete
          </div>
          
          <h1 className="text-5xl md:text-7xl font-extrabold text-transparent bg-clip-text bg-gradient-to-b from-white to-zinc-500 tracking-tight mb-8">
            Master Game Engine <br className="hidden md:block" /> Algorithms.
          </h1>
          
          <p className="text-lg md:text-xl text-zinc-400 mb-10 max-w-2xl leading-relaxed">
            The ultimate platform to level up your game development math and logic skills. Solve real-world game engine problems, optimize your code, and rise up the global leaderboard.
          </p>
          
          <div className="flex items-center gap-4">
            <Link 
              href="/problems" 
              className="px-8 py-4 rounded-full bg-indigo-600 text-white font-semibold hover:bg-indigo-500 transition-all shadow-[0_0_40px_-10px_rgba(79,70,229,0.5)] hover:shadow-[0_0_60px_-15px_rgba(79,70,229,0.7)] hover:-translate-y-0.5"
            >
              Start Coding Now
            </Link>
            <a 
              href="https://github.com"
              target="_blank"
              rel="noreferrer"
              className="px-8 py-4 rounded-full border border-white/10 bg-white/5 text-white font-semibold hover:bg-white/10 transition-all"
            >
              View on GitHub
            </a>
          </div>
        </div>
      </main>

      {/* Decorative Grid */}
      <div className="absolute inset-0 bg-[url('https://grainy-gradients.vercel.app/noise.svg')] opacity-20 mix-blend-overlay pointer-events-none" />
    </div>
  );
}
