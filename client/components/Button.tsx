import { cn } from "@/lib/utlis";
import React from "react";
import {
  TextProps,
  TouchableOpacity,
  TouchableOpacityProps,
} from "react-native";
import { Text } from "./Themed";

const Button = (
  props: TouchableOpacityProps & {
    children?: React.ReactNode;
    className?: string;
  },
) => {
  const { className, children, ...otherProps } = props;

  return (
    <TouchableOpacity
      className={cn("bg-blue-500 !text-white p-3 rounded-md", className)}
      {...otherProps}
    >
      {children}
    </TouchableOpacity>
  );
};

export default Button;

export const ButtonText = (props: TextProps) => {
  const { children, ...otherProps } = props;
  return (
    <Text {...otherProps} className={cn(otherProps.className, "text-white")}>
      {children}
    </Text>
  );
};
