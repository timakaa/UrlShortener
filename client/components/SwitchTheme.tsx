import AsyncStorage from "@react-native-async-storage/async-storage";
import { Text, View } from "@/components/Themed";
import { Switch } from "react-native";
import { useColorScheme } from "nativewind";

const SwitchTheme = () => {
  const { colorScheme, toggleColorScheme } = useColorScheme();

  return (
    <View className='flex-row items-center justify-between w-full'>
      <Text className='text-xl dark:text-white font-bold'>Dark Mode</Text>
      <Switch
        value={colorScheme === "dark"}
        onValueChange={(value) => {
          setTimeout(() => {
            AsyncStorage.setItem("theme", value ? "dark" : "light")
              .then(() => toggleColorScheme())
              .catch((error) =>
                console.error("Ошибка при сохранении темы:", error),
              );
          }, 100);
        }}
      />
    </View>
  );
};

export default SwitchTheme;
