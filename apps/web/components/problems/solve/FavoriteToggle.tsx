import React from "react";
import { Heart } from "lucide-react";
import { useMutation } from "@tanstack/react-query";
import { api } from "@/lib/api";
import { clsx } from "clsx";

interface FavoriteToggleProps {
  slug: string;
  initialIsFavorite?: boolean;
}

export function FavoriteToggle({ slug, initialIsFavorite = false }: FavoriteToggleProps) {
  const [isFavorite, setIsFavorite] = React.useState(initialIsFavorite);

  const mutation = useMutation({
    mutationFn: () => api.post(`problems/${slug}/favorite`),
    onMutate: () => {
      setIsFavorite((prev) => !prev);
    },
    onError: () => {
      // Revert on error
      setIsFavorite((prev) => !prev);
    },
  });

  return (
    <button
      onClick={() => mutation.mutate()}
      className="p-1.5 transition-colors hover:text-red-500 disabled:opacity-50"
      disabled={mutation.isPending}
      title={isFavorite ? "Remove from favorites" : "Add to favorites"}
    >
      <Heart
        className={clsx("h-5 w-5", {
          "fill-red-500 text-red-500": isFavorite,
          "text-zinc-400": !isFavorite,
        })}
      />
    </button>
  );
}
