import React from "react";
import FontAwesome from "@expo/vector-icons/FontAwesome";
import { Tabs } from "expo-router";
import { StatusBar } from "expo-status-bar";
import { Platform, Pressable } from "react-native";

import Colors from "@/constants/Colors";
import { useColorScheme } from "@/components/useColorScheme";
import { useClientOnlyValue } from "@/components/useClientOnlyValue";

// You can explore the built-in icon families and icons on the web at https://icons.expo.fyi/
function TabBarIcon(props: {
  name: React.ComponentProps<typeof FontAwesome>["name"];
  color: string;
}) {
  return <FontAwesome size={28} style={{ marginBottom: -3 }} {...props} />;
}

export default function TabLayout() {
  const colorScheme = useColorScheme();

  return (
    <>
      <StatusBar />
      <Tabs
        screenOptions={{
          tabBarActiveTintColor: Colors[colorScheme ?? "light"].tint,
          // Disable the static render of the header on web
          // to prevent a hydration error in React Navigation v6.
          headerShown: useClientOnlyValue(false, true),
          tabBarStyle: Platform.select({
            ios: {
              position: "absolute",
              height: 95,
              paddingTop: 10,
              paddingBottom: 10,
            },
            default: {
              marginTop: 0,
              height: 60,
              paddingBottom: 10,
            },
          }),
          tabBarButton: (props) => (
            <Pressable {...props} hitSlop={{ top: 10 }} />
          ),
        }}
      >
        <Tabs.Screen
          name='index'
          options={{
            title: "Weather",
            tabBarIcon: ({ color }) => (
              <TabBarIcon name='cloud' color={color} />
            ),
          }}
        />
        <Tabs.Screen
          name='settings'
          options={{
            title: "Settings",
            tabBarIcon: ({ color }) => <TabBarIcon name='cog' color={color} />,
          }}
        />
      </Tabs>
    </>
  );
}
