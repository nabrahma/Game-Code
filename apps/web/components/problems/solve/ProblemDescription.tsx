import React from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import rehypeRaw from "rehype-raw";
import clsx from "clsx";

interface ProblemDescriptionProps {
  content: string;
  className?: string;
}

export function ProblemDescription({ content, className }: ProblemDescriptionProps) {
  return (
    <div className={clsx("prose prose-invert max-w-none prose-pre:bg-zinc-900 prose-pre:border prose-pre:border-zinc-800", className)}>
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[rehypeRaw]}
        components={{
          code({ node, inline, className, children, ...props }: any) {
            const match = /language-(\w+)/.exec(className || "");
            return !inline ? (
              <code className={className} {...props}>
                {children}
              </code>
            ) : (
              <code className="rounded bg-zinc-800 px-1 py-0.5 text-sm text-zinc-200 font-mono" {...props}>
                {children}
              </code>
            );
          },
        }}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
}
