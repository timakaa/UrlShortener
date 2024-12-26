import { Text, View } from "@/components/Themed";
import { IUrl } from "@/interfaces/Url";
import React from "react";
import { useUrls } from "@/context/UrlsContext";
import Url from "./Url";

const Urls = () => {
  const { urls: data, isLoading, error } = useUrls();

  if (isLoading) return <Text>Loading...</Text>;
  if (error) return <Text>Error: {error.message}</Text>;
  if (!data || data.length === 0) return <Text>No URLs found</Text>;

  return (
    <View className='flex-row flex-wrap justify-between'>
      {data.map((url: IUrl) => (
        <Url key={url.ID} url={url} />
      ))}
    </View>
  );
};

export default Urls;
