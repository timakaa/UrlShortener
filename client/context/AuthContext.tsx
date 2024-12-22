import React, { createContext, useState, useContext, useEffect } from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";

type AuthContextType = {
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (token: string) => void;
  logout: () => void;
  token: string | null;
};

const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [token, setToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    loadToken();
  }, []);

  const loadToken = async () => {
    try {
      const savedToken = await AsyncStorage.getItem("token");
      if (savedToken) {
        setToken(savedToken);
        router.push("/(tabs)");
      }
    } catch (error) {
      console.error("Error loading token:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const login = async (newToken: string) => {
    try {
      await AsyncStorage.setItem("token", newToken);
      setToken(newToken);
      router.push("/(tabs)");
    } catch (error) {
      console.error("Error saving token:", error);
    }
  };

  const logout = async () => {
    try {
      await AsyncStorage.removeItem("token");
      setToken(null);
    } catch (error) {
      console.error("Error removing token:", error);
    }
  };

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated: !!token,
        isLoading,
        login,
        logout,
        token,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
