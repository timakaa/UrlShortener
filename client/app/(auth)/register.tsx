import { useState } from "react";
import { View, Text, TextInput } from "@/components/Themed";
import { useAuth } from "../../context/AuthContext";
import { Link, router } from "expo-router";
import { API_URL } from "@env";
import Button from "@/components/Button";

export default function RegisterScreen() {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    try {
      const response = await fetch(`${API_URL}/register/init`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password, username }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        console.error("Server error:", errorText);
        return;
      }

      router.push({
        pathname: "/(auth)/verify",
        params: { email },
      });
    } catch (error) {
      console.error("Register error:", error);
    }
  };

  return (
    <View className='flex-1 justify-center p-4 gap-y-4'>
      <TextInput
        placeholder='Username'
        value={username}
        onChangeText={setUsername}
      />
      <TextInput placeholder='Email' value={email} onChangeText={setEmail} />
      <TextInput
        placeholder='Password'
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      <Button onPress={handleLogin}>
        <Text className='text-lg font-bold text-center'>Sign up</Text>
      </Button>
      <View className='justify-between px-4 flex-row items-center'>
        <Text className='text-lg text-center'>Already have an account?</Text>
        <Link href='/(auth)/login'>
          <Text className='text-lg text-center text-blue-500 dark:text-blue-500'>
            Sign in
          </Text>
        </Link>
      </View>
    </View>
  );
}
