import React from "react";
import { Text, View } from "@/components/Themed";
import Button, { ButtonText } from "@/components/Button";
import { useAuth } from "@/context/AuthContext";
import { MaterialIcons } from "@expo/vector-icons";

const Profile = () => {
  const { logout } = useAuth();

  return (
    <View className='flex-1'>
      <Button
        className='w-1/2 justify-between flex-row items-center mt-10 mx-auto'
        onPress={logout}
      >
        <ButtonText>Logout</ButtonText>
        <MaterialIcons name='exit-to-app' size={20} color='white' />
      </Button>
    </View>
  );
};

export default Profile;
