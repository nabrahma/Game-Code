import React from "react";

export default function SolveLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="h-[calc(100vh-64px)] w-full overflow-hidden bg-[#0A0A0A] p-2">
      {children}
    </div>
  );
}
