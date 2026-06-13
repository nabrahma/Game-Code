import { useState, useCallback, useRef, useEffect } from "react";
import { api } from "../api";

interface RunResult {
  runId: string;
}

export interface StreamEvent {
  run_id: string;
  status: "queued" | "running" | "success" | "error" | "timeout";
  output?: string;
}

export function useExecution() {
  const [isExecuting, setIsExecuting] = useState(false);
  const [output, setOutput] = useState<string>("");
  const [status, setStatus] = useState<string>("");
  const eventSourceRef = useRef<EventSource | null>(null);

  const cleanup = useCallback(() => {
    if (eventSourceRef.current) {
      eventSourceRef.current.close();
      eventSourceRef.current = null;
    }
  }, []);

  // Cleanup on unmount
  useEffect(() => {
    return cleanup;
  }, [cleanup]);

  const executeCode = useCallback(async (
    problemId: string, 
    language: string, 
    code: string, 
    input?: string
  ) => {
    setIsExecuting(true);
    setOutput("");
    setStatus("submitting");

    try {
      // 1. Enqueue job
      const result = await api.post("run", {
        json: {
          problem_id: problemId,
          language,
          code,
          input,
        }
      }).json<RunResult>();

      // 2. Connect to SSE stream
      setStatus("connecting...");
      const url = `${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'}/run/${result.runId}/stream`;
      
      const sse = new EventSource(url);
      eventSourceRef.current = sse;

      sse.onmessage = (event) => {
        try {
          const data: StreamEvent = JSON.parse(event.data);
          setStatus(data.status);
          if (data.output) {
            setOutput((prev) => prev + data.output + "\n");
          }
          if (["success", "error", "timeout"].includes(data.status)) {
            setIsExecuting(false);
            cleanup();
          }
        } catch (e) {
          console.error("Failed to parse SSE event", e);
        }
      };

      sse.onerror = () => {
        setStatus("connection lost");
        setIsExecuting(false);
        cleanup();
      };

    } catch (err) {
      console.error(err);
      setOutput("Failed to enqueue execution. Check API connection.");
      setStatus("error");
      setIsExecuting(false);
    }
  }, [cleanup]);

  return { executeCode, isExecuting, output, status, setOutput };
}
