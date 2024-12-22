import { useState } from "react";
import { TextInput } from "react-native";
import { View, Text } from "@/components/Themed";
import { CodeField, Cursor } from "react-native-confirmation-code-field";
import { router, useLocalSearchParams } from "expo-router";
import { useAuth } from "@/context/AuthContext";
import { API_URL } from "@env";

const CELL_COUNT = 6;

const Verify = () => {
  const [value, setValue] = useState("");
  const { email } = useLocalSearchParams();
  const { login } = useAuth();

  const handleCodeChange = async (code: string) => {
    setValue(code);

    if (code.length === CELL_COUNT) {
      console.log("Verification code entered:", code);
      const response = await fetch(`${API_URL}/register/verify`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ code, email }),
      });

      if (response.ok) {
        const data = await response.json();
        login(data.token);
        router.replace("/(tabs)");
      } else {
        console.error("Failed to verify code");
      }
    }
  };

  return (
    <View className='flex-1 items-center justify-center p-5'>
      <Text className='text-2xl font-bold mb-8'>Confirm your email</Text>
      {email && (
        <Text className='text-gray-500 mb-4'>
          We sent a verification code to {email as string}
        </Text>
      )}
      <CodeField
        value={value}
        onChangeText={handleCodeChange}
        cellCount={CELL_COUNT}
        InputComponent={TextInput}
        rootStyle={{ marginTop: 20 }}
        keyboardType='number-pad'
        textContentType='oneTimeCode'
        renderCell={({ index, symbol, isFocused }) => (
          <View
            key={index}
            className={`w-12 h-12 border-2 rounded-lg m-0.5 items-center justify-center
              ${isFocused ? "border-blue-500" : "border-gray-200"}`}
          >
            <Text className='text-2xl text-center'>
              {symbol || (isFocused ? <Cursor /> : null)}
            </Text>
          </View>
        )}
      />
    </View>
  );
};

export default Verify;
