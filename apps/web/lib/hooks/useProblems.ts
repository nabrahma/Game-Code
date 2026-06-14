import { useQuery } from "@tanstack/react-query";
import { api } from "../api";

export interface Tag {
  id: string;
  name: string;
  slug: string;
  category?: string;
}

export interface ProblemSummary {
  id: string;
  slug: string;
  title: string;
  difficulty: "easy" | "medium" | "hard";
  acceptance_rate: number;
  tags: Tag[];
  user_status?: "attempted" | "solved";
  is_favorite?: boolean;
}

export interface ProblemListResponse {
  items: ProblemSummary[];
  total: number;
}

export interface ProblemFilter {
  difficulty?: string;
  q?: string;
  sort?: string;
  offset?: number;
  limit?: number;
}

export function useProblems(filter: ProblemFilter) {
  return useQuery({
    queryKey: ["problems", filter],
    queryFn: async () => {
      const searchParams = new URLSearchParams();
      if (filter.difficulty) searchParams.set("difficulty", filter.difficulty);
      if (filter.q) searchParams.set("q", filter.q);
      if (filter.sort) searchParams.set("sort", filter.sort);
      if (filter.offset) searchParams.set("offset", filter.offset.toString());
      if (filter.limit) searchParams.set("limit", filter.limit.toString());

      const res = await api.get(`problems?${searchParams.toString()}`).json<ProblemListResponse>();
      return res;
    },
    staleTime: 1000 * 30, // 30 seconds to match API cache TTL
  });
}

export function useProblem(slug: string) {
  return useQuery({
    queryKey: ["problem", slug],
    queryFn: async () => {
      return await api.get(`problems/${slug}`).json<any>(); // any for now until we type Problem detail
    },
    enabled: !!slug,
  });
}
