import React from "react";
import { Text, View } from "@/components/Themed";
import Button from "@/components/Button";
import { useAuth } from "@/context/AuthContext";

const Profile = () => {
  const { logout } = useAuth();

  return (
    <View className='flex-1'>
      <Text>Profile</Text>
      <Button onPress={logout}>
        <Text>Logout</Text>
      </Button>
    </View>
  );
};

export default Profile;
