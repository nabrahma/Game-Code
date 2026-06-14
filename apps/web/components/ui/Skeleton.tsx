import React from 'react';
import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function Skeleton({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn("animate-pulse rounded-md bg-zinc-800", className)}
      {...props}
    />
  );
}

// Pre-built layout skeletons for convenience
export function TableSkeleton({ rows = 5 }: { rows?: number }) {
  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-6">
        <Skeleton className="h-8 w-1/4" />
        <Skeleton className="h-10 w-32 rounded-lg" />
      </div>
      <div className="border border-zinc-800 rounded-xl overflow-hidden">
        <div className="bg-zinc-950 px-6 py-4 border-b border-zinc-800">
          <Skeleton className="h-4 w-full" />
        </div>
        <div className="divide-y divide-zinc-800/50">
          {Array.from({ length: rows }).map((_, i) => (
            <div key={i} className="px-6 py-4 flex items-center justify-between">
              <div className="space-y-2 w-1/3">
                <Skeleton className="h-5 w-3/4" />
                <Skeleton className="h-3 w-1/2" />
              </div>
              <Skeleton className="h-6 w-16 rounded-full" />
              <Skeleton className="h-5 w-12" />
              <div className="flex gap-2">
                <Skeleton className="h-8 w-8 rounded" />
                <Skeleton className="h-8 w-8 rounded" />
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
