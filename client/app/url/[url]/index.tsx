import React, { useEffect } from "react";
import { useLocalSearchParams } from "expo-router";
import { useGetUrl } from "@/hooks/url.hooks";
import { Text } from "@/components/Themed";

const RedirectUrlPage = () => {
  const { url } = useLocalSearchParams();
  const { data, isLoading } = useGetUrl(url as string);
  useEffect(() => {
    if (data) {
      window.location.href = data.originalUrl;
    }
  }, [data]);

  if (isLoading) {
    return <Text>Loading...</Text>;
  }

  return <></>;
};

export default RedirectUrlPage;
