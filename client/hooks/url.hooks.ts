import { API_URL } from "@env";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "../context/AuthContext";

export const useCreateUrl = () => {
  const { token } = useAuth();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (url: string) => {
      const response = await fetch(`${API_URL}/url`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ original_url: url }),
      });

      if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || "Failed to create URL");
      }

      return response.json();
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["urls"] });
    },
  });
};

interface ErrorResponse {
  error: string;
}

export const useGetUrls = () => {
  const { token } = useAuth();

  return useQuery({
    queryKey: ["urls"],
    queryFn: async () => {
      try {
        const url = `${API_URL}/urls`;
        console.log("Making request to:", url);

        const response = await fetch(url, {
          method: "GET",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        });

        const data = await response.json();

        // If response is not ok, throw error with server message
        if (!response.ok) {
          const errorData = data as ErrorResponse;
          throw new Error(errorData.error || "Failed to fetch URLs");
        }

        return data;
      } catch (error) {
        if (error instanceof Error) {
          throw error;
        }
        throw new Error("An unknown error occurred");
      }
    },
    enabled: !!token,
    retry: false,
  });
};

export const useGetUrl = (shortUrl: string) => {
  return useQuery({
    queryKey: ["url", shortUrl],
    queryFn: async () => {
      const response = await fetch(`${API_URL}/url?code=${shortUrl}`);
      const data = await response.json();
      return data;
    },
  });
};
