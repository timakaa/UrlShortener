import React, { useEffect } from "react";
import FontAwesome from "@expo/vector-icons/FontAwesome";
import { Tabs, Redirect } from "expo-router";
import { StatusBar } from "expo-status-bar";
import {
  Platform,
  Pressable,
  StyleProp,
  View,
  ViewStyle,
  ActivityIndicator,
} from "react-native";

import Colors from "@/constants/Colors";
import { useColorScheme } from "@/components/useColorScheme";
import { useClientOnlyValue } from "@/components/useClientOnlyValue";
import { useAuth } from "@/context/AuthContext";

// You can explore the built-in icon families and icons on the web at https://icons.expo.fyi/
function TabBarIcon(props: {
  name: React.ComponentProps<typeof FontAwesome>["name"];
  color: string;
  className?: string;
  style?: StyleProp<ViewStyle>;
}) {
  return <FontAwesome size={28} style={{ marginBottom: -3 }} {...props} />;
}

export default function TabLayout() {
  const { isAuthenticated, isLoading } = useAuth();
  const colorScheme = useColorScheme();

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size='large' />
      </View>
    );
  }

  if (!isAuthenticated) {
    return <Redirect href='/(auth)/login' />;
  }

  return (
    <>
      <StatusBar />
      <Tabs
        screenOptions={{
          tabBarActiveTintColor: Colors[colorScheme ?? "light"].tint,
          headerShown: useClientOnlyValue(false, true),
          tabBarIconStyle: {
            marginBottom: 4,
          },
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
            title: "Url",
            tabBarIcon: ({ color }) => (
              <View className='rotate-[135deg]'>
                <TabBarIcon name='link' color={color} />
              </View>
            ),
          }}
        />
        <Tabs.Screen
          name='profile'
          options={{
            title: "Profile",
            tabBarIcon: ({ color }) => (
              <View>
                <TabBarIcon name='user-circle' color={color} />
              </View>
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
