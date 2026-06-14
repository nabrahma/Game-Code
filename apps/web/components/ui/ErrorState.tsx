import React from 'react';
import { AlertCircle, RefreshCcw } from 'lucide-react';

interface ErrorStateProps {
  title?: string;
  message?: string;
  onRetry?: () => void;
}

export function ErrorState({ 
  title = "Something went wrong", 
  message = "An unexpected error occurred while loading this data.",
  onRetry 
}: ErrorStateProps) {
  return (
    <div className="flex flex-col items-center justify-center p-12 text-center border border-rose-900/30 rounded-xl bg-rose-500/5">
      <AlertCircle className="w-12 h-12 text-rose-500 mb-4" />
      <h3 className="text-lg font-medium text-rose-100">{title}</h3>
      <p className="text-rose-400/80 mt-2 max-w-sm">{message}</p>
      
      {onRetry && (
        <button 
          onClick={onRetry}
          className="mt-6 flex items-center gap-2 bg-rose-500/10 hover:bg-rose-500/20 text-rose-400 px-4 py-2 rounded-lg font-medium transition-colors border border-rose-500/20"
        >
          <RefreshCcw className="w-4 h-4" />
          Try Again
        </button>
      )}
    </div>
  );
}
