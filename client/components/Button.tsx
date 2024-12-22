import React from "react";
import { TouchableOpacity, TouchableOpacityProps } from "react-native";

const Button = (
  props: TouchableOpacityProps & {
    children?: React.ReactNode;
    className?: string;
  },
) => {
  const { className, children, ...otherProps } = props;

  return (
    <TouchableOpacity
      className={`bg-blue-500 p-3 rounded-md ${className || ""}`}
      {...otherProps}
    >
      {children}
    </TouchableOpacity>
  );
};

export default Button;
