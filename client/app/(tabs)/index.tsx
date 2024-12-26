import { View, Text, TextInput } from "@/components/Themed";
import Button from "@/components/Button";
import Urls from "@/components/Urls";
import { useCreateUrl } from "@/hooks/url.hooks";
import { useState } from "react";
import { RefreshControl, ScrollView } from "react-native";
import { UrlsProvider, useUrls } from "@/context/UrlsContext";
import React from "react";

export default function WeatherScreen() {
  return (
    <UrlsProvider>
      <ScreenWithProvider />
    </UrlsProvider>
  );
}

function ScreenWithProvider() {
  const { mutate: createUrl } = useCreateUrl();
  const [url, setUrl] = useState("");
  const { refetch } = useUrls();
  const [refreshing, setRefreshing] = useState(false);

  const onRefresh = React.useCallback(() => {
    setRefreshing(true);
    refetch().finally(() => setRefreshing(false));
  }, [refetch]);

  return (
    <ScrollView
      className='p-5'
      refreshControl={
        <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
      }
    >
      <View className='flex-row w-full items-center gap-x-4'>
        <TextInput
          className='text-xl font-bold flex-1'
          placeholder='Url Shortener'
          value={url}
          onChangeText={setUrl}
        />
        <Button
          className='bg-blue-500 text-white py-3 px-6 rounded-md'
          onPress={() => {
            createUrl(url);
          }}
        >
          <Text className='text-lg font-bold'>Create</Text>
        </Button>
      </View>
      <View className='my-10 flex-row items-center gap-x-4'>
        <Urls />
      </View>
    </ScrollView>
  );
}
