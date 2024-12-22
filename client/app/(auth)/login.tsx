import { useState } from "react";
import { View, Text, TextInput } from "@/components/Themed";
import { useAuth } from "../../context/AuthContext";
import { Link, router } from "expo-router";
import { API_URL } from "@env";
import Button from "@/components/Button";

export default function LoginScreen() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { login } = useAuth();

  const handleLogin = async () => {
    try {
      console.log(API_URL);
      const response = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        console.error("Server error:", errorText);
        return;
      }

      const data = await response.json();

      if (data.token) {
        login(data.token);
      } else {
        console.error("No token received:", data);
      }
    } catch (error) {
      console.error("Login error:", error);
    }
  };

  return (
    <View className='flex-1 justify-center p-4 gap-y-4'>
      <TextInput placeholder='Email' value={email} onChangeText={setEmail} />
      <TextInput
        placeholder='Password'
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      <Button onPress={handleLogin}>
        <Text className='text-lg font-bold text-center'>Sign in</Text>
      </Button>
      <View className='justify-between px-4 flex-row items-center'>
        <Text className='text-lg text-center'>Don't have an account?</Text>
        <Link href='/(auth)/register'>
          <Text className='text-lg text-center text-blue-500 dark:text-blue-500'>
            Sign up
          </Text>
        </Link>
      </View>
    </View>
  );
}
