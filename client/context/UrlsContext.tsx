import { createContext, useContext } from "react";
import { Url } from "@/interfaces/Url";
import { useGetUrls } from "@/hooks/url.hooks";

type UrlsContextType = {
  urls: Url[];
  isLoading: boolean;
  error: Error | null;
  refetch: () => Promise<unknown>;
};

const UrlsContext = createContext<UrlsContextType | null>(null);

export function UrlsProvider({ children }: { children: React.ReactNode }) {
  const { data, isLoading, error, refetch } = useGetUrls();

  return (
    <UrlsContext.Provider value={{ urls: data, isLoading, error, refetch }}>
      {children}
    </UrlsContext.Provider>
  );
}

export const useUrls = () => {
  const context = useContext(UrlsContext);
  if (!context) {
    throw new Error("useUrls must be used within an UrlsProvider");
  }
  return context;
};
