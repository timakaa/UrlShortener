import { Stack } from "expo-router";

export default function AuthLayout() {
  return (
    <Stack>
      <Stack.Screen
        name='login'
        options={{
          title: "Login",
          headerShown: true,
          headerLeft: () => null,
          headerBackVisible: false,
        }}
      />
      <Stack.Screen
        name='register'
        options={{
          title: "Register",
          headerShown: true,
          headerLeft: () => null,
          headerBackVisible: false,
        }}
      />
      <Stack.Screen
        name='verify'
        options={{
          title: "Verify",
          headerShown: true,
        }}
      />
    </Stack>
  );
}
