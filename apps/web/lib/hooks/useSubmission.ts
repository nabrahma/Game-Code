import { useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "../api";
import { useState, useEffect } from "react";

export type SubmissionVerdict = "pending" | "accepted" | "wrong_answer" | "time_limit_exceeded" | "memory_limit_exceeded" | "runtime_error" | "compile_error" | "internal_error";

export interface TestResult {
  id: string;
  input: string;
  expected: string;
  actual: string;
  passed: boolean;
  runtime_ms?: number;
  memory_kb?: number;
}

export interface Submission {
  id: string;
  problem_id: string;
  language: string;
  code: string;
  verdict: SubmissionVerdict;
  runtime_ms?: number;
  memory_kb?: number;
  error_message?: string;
  passed_test_count?: number;
  total_test_count?: number;
  results?: TestResult[];
  created_at: string;
}

export function useSubmission() {
  const queryClient = useQueryClient();
  const [activeSubmissionId, setActiveSubmissionId] = useState<string | null>(null);
  const [submissionData, setSubmissionData] = useState<Submission | null>(null);

  // Poll for active submission
  useEffect(() => {
    if (!activeSubmissionId) return;

    const interval = setInterval(async () => {
      try {
        const sub = await api.get(`submissions/${activeSubmissionId}`).json<Submission>();
        setSubmissionData(sub);

        if (sub.verdict !== "pending") {
          setActiveSubmissionId(null); // Stop polling
          queryClient.invalidateQueries({ queryKey: ["problems"] }); // Refresh progress
        }
      } catch (err) {
        console.error("Polling error:", err);
      }
    }, 1500);

    return () => clearInterval(interval);
  }, [activeSubmissionId, queryClient]);

  const submitMutation = useMutation({
    mutationFn: async ({ problem_id, language, code }: { problem_id: string; language: string; code: string }) => {
      const res = await api.post("submissions", { json: { problem_id, language, code } }).json<Submission>();
      return res;
    },
    onSuccess: (data) => {
      setSubmissionData(data);
      setActiveSubmissionId(data.id);
    },
  });

  return {
    submit: submitMutation.mutate,
    isSubmitting: submitMutation.isPending || !!activeSubmissionId,
    submission: submissionData,
  };
}
