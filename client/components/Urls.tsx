import { useGetUrls } from "@/hooks/url.hooks";
import { Text, View } from "@/components/Themed";
import { Url } from "@/interfaces/Url";
import {
  Linking,
  TouchableOpacity,
  ScrollView,
  RefreshControl,
} from "react-native";
import React from "react";
import { useUrls } from "@/context/UrlsContext";

const Urls = () => {
  const { urls: data, isLoading, error, refetch } = useUrls();
  const [refreshing, setRefreshing] = React.useState(false);

  const onRefresh = React.useCallback(() => {
    setRefreshing(true);
    refetch().finally(() => setRefreshing(false));
  }, [refetch]);

  if (isLoading) return <Text>Loading...</Text>;
  if (error) return <Text>Error: {error.message}</Text>;
  if (!data || data.length === 0) return <Text>No URLs found</Text>;

  return (
    <ScrollView
      refreshControl={
        <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
      }
    >
      <View className='flex-row flex-wrap justify-between'>
        {data.map((url: Url) => (
          <TouchableOpacity
            key={url.ID}
            onPress={() =>
              url.original_url && Linking.openURL(url.original_url)
            }
            className='bg-blue-500 text-white w-[48%] mb-4 py-3 px-6 rounded-md'
          >
            <Text>{url.original_url.split("://")[1]}</Text>
          </TouchableOpacity>
        ))}
      </View>
    </ScrollView>
  );
};

export default Urls;
