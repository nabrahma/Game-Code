import React from 'react';
import { FileQuestion } from 'lucide-react';

interface EmptyStateProps {
  title?: string;
  message?: string;
  icon?: React.ReactNode;
  action?: React.ReactNode;
}

export function EmptyState({ 
  title = "No results found", 
  message = "Try adjusting your filters or search query.", 
  icon = <FileQuestion className="w-12 h-12 text-zinc-600 mb-4" />,
  action 
}: EmptyStateProps) {
  return (
    <div className="flex flex-col items-center justify-center p-12 text-center border border-dashed border-zinc-800 rounded-xl bg-zinc-900/30">
      {icon}
      <h3 className="text-lg font-medium text-zinc-200">{title}</h3>
      <p className="text-zinc-500 mt-2 max-w-sm">{message}</p>
      {action && <div className="mt-6">{action}</div>}
    </div>
  );
}
