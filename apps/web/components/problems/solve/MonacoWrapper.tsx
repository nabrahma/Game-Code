import React, { useEffect } from "react";
import Editor, { useMonaco } from "@monaco-editor/react";

interface MonacoWrapperProps {
  language: string;
  code: string;
  onChange: (value: string | undefined) => void;
}

export function MonacoWrapper({ language, code, onChange }: MonacoWrapperProps) {
  const monaco = useMonaco();

  useEffect(() => {
    if (monaco) {
      monaco.editor.defineTheme("gamecode-dark", {
        base: "vs-dark",
        inherit: true,
        rules: [],
        colors: {
          "editor.background": "#09090b", // zinc-950
          "editor.lineHighlightBackground": "#18181b", // zinc-900
        },
      });
      monaco.editor.setTheme("gamecode-dark");
    }
  }, [monaco]);

  const mapLanguage = (lang: string) => {
    return lang;
  };

  return (
    <div className="h-full w-full bg-zinc-950 pt-2">
      <Editor
        height="100%"
        language={mapLanguage(language)}
        value={code}
        onChange={onChange}
        theme="gamecode-dark"
        options={{
          minimap: { enabled: false },
          fontSize: 14,
          fontFamily: "'JetBrains Mono', 'Fira Code', monospace",
          wordWrap: "on",
          padding: { top: 16 },
          scrollBeyondLastLine: false,
          smoothScrolling: true,
          cursorBlinking: "smooth",
          cursorSmoothCaretAnimation: "on",
        }}
        loading={
          <div className="flex h-full items-center justify-center text-sm text-zinc-500">
            Loading editor...
          </div>
        }
      />
    </div>
  );
}
