import { View, Text } from "@/components/Themed";

export default function WeatherScreen() {
  return (
    <View className='flex-1 items-center justify-center'>
      <Text className='text-2xl font-bold'>Weather</Text>
      <View className='my-10 flex-row items-center gap-x-4'></View>
    </View>
  );
}
