import React from "react";
import { TouchableOpacity, Linking, View } from "react-native";
import { Text } from "@/components/Themed";
import { IUrl } from "@/interfaces/Url";
import * as Clipboard from "expo-clipboard";
import { CLIENT_URL } from "@env";
import Button, { ButtonText } from "./Button";

const Url = ({ url }: { url: IUrl }) => {
  return (
    <View className='flex-col justify-between mb-4 w-[48.5%] border border-gray-700 rounded-md p-2 items-center'>
      <View className='mb-4 w-full'>
        <View className='flex-row items-start justify-start w-full'>
          <Text className='text-gray-500'>Short URL: </Text>
          <Text className='text-gray-500'>{url.short_url}</Text>
        </View>
        <View className='flex-row items-start justify-start w-full'>
          <Text className='text-gray-500'>Views: </Text>
          <Text className='text-gray-500'>{url.visits}</Text>
        </View>
      </View>
      <Button
        onPress={() => url.original_url && Linking.openURL(url.original_url)}
        className='bg-blue-500 mb-2 text-white py-3 w-full px-4 rounded-md'
      >
        <ButtonText
          numberOfLines={1}
          ellipsizeMode='tail'
          className='w-full text-start'
        >
          {url.original_url.split("://")[1]}
        </ButtonText>
      </Button>
      <Button
        onPress={() => {
          if (url.short_url) {
            Clipboard.setStringAsync(`${CLIENT_URL}/url/${url.short_url}`);
          }
        }}
        className='bg-blue-500 text-white py-3 w-full px-4 rounded-md'
      >
        <ButtonText>Copy</ButtonText>
      </Button>
    </View>
  );
};

export default Url;
