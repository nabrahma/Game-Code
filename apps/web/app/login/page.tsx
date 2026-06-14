"use client";

import { useAuth } from "@/lib/auth";
import { redirect } from "next/navigation";
import { useState } from "react";

export default function LoginPage() {
  const { user, isLoading } = useAuth();
  const [email, setEmail] = useState("");
  const [magicLinkSent, setMagicLinkSent] = useState(false);

  if (isLoading) return <div className="flex min-h-screen items-center justify-center">Loading...</div>;
  if (user) {
    redirect("/");
  }

  const handleMagicLink = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email) return;
    try {
      // We will implement the actual API call later
      // await api.post("auth/magic-link", { json: { email } });
      setMagicLinkSent(true);
    } catch (err) {
      console.error(err);
    }
  };

  const handleOAuth = (provider: "github" | "google") => {
    window.location.href = `${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'}/auth/oauth/${provider}`;
  };

  return (
    <div className="flex min-h-screen items-center justify-center p-4">
      <div className="w-full max-w-sm rounded-xl border border-zinc-800 bg-zinc-900 p-8 shadow-2xl">
        <div className="mb-8 text-center">
          <h1 className="text-2xl font-bold tracking-tight text-white">Welcome back</h1>
          <p className="text-sm text-zinc-400 mt-2">Sign in to your account</p>
        </div>

        <div className="space-y-4">
          <button
            onClick={() => handleOAuth("github")}
            className="flex w-full items-center justify-center gap-3 rounded-md bg-[#24292F] px-4 py-2.5 text-sm font-semibold text-white transition-colors hover:bg-[#24292F]/90"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="h-5 w-5"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"/><path d="M9 18c-4.51 2-5-2-7-2"/></svg>
            Continue with GitHub
          </button>
          
          {/* We will add Google OAuth later since it's in the PRD, but let's stick to GitHub for now to match UI or we can add both */}
          <button
            onClick={() => handleOAuth("google")}
            className="flex w-full items-center justify-center gap-3 rounded-md bg-white px-4 py-2.5 text-sm font-semibold text-zinc-900 transition-colors hover:bg-zinc-100"
          >
            <svg className="h-5 w-5" viewBox="0 0 24 24">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
            </svg>
            Continue with Google
          </button>

          <div className="relative my-6">
            <div className="absolute inset-0 flex items-center">
              <div className="w-full border-t border-zinc-800"></div>
            </div>
            <div className="relative flex justify-center text-sm">
              <span className="bg-zinc-900 px-2 text-zinc-500">Or continue with email</span>
            </div>
          </div>

          {magicLinkSent ? (
            <div className="rounded-md bg-green-500/10 p-4 text-sm text-green-400">
              Check your email for a login link!
            </div>
          ) : (
            <form onSubmit={handleMagicLink} className="space-y-4">
              <div>
                <label htmlFor="email" className="sr-only">Email address</label>
                <div className="relative">
                  <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="h-5 w-5 text-zinc-500"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></svg>
                  </div>
                  <input
                    id="email"
                    type="email"
                    required
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="block w-full rounded-md border border-zinc-800 bg-zinc-950 py-2.5 pl-10 pr-3 text-sm text-white placeholder-zinc-500 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
                    placeholder="you@example.com"
                  />
                </div>
              </div>
              <button
                type="submit"
                className="w-full rounded-md bg-indigo-600 px-4 py-2.5 text-sm font-semibold text-white transition-colors hover:bg-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-zinc-900"
              >
                Send Magic Link
              </button>
            </form>
          )}
        </div>
      </div>
    </div>
  );
}
