import SwitchTheme from "@/components/SwitchTheme";
import { View } from "@/components/Themed";

export default function TabTwoScreen() {
  return (
    <View className='flex-1 items-start justify-start px-6 pt-6'>
      <SwitchTheme />
    </View>
  );
}
