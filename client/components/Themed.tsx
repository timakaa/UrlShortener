/**
 * Learn more about Light and Dark modes:
 * https://docs.expo.io/guides/color-schemes/
 */

import { cn } from "@/lib/utlis";
import React from "react";
import {
  Text as DefaultText,
  View as DefaultView,
  TextInput as DefaultTextInput,
  TextInputProps,
} from "react-native";

type ThemeProps = {
  className?: string;
};

export type TextProps = ThemeProps & DefaultText["props"];
export type ViewProps = ThemeProps & DefaultView["props"];

export function Text(props: TextProps) {
  const { className, ...otherProps } = props;

  return (
    <DefaultText
      className={cn("text-black dark:text-white", className)}
      {...otherProps}
    />
  );
}

export function View(props: ViewProps) {
  const { className, ...otherProps } = props;

  return (
    <DefaultView
      className={cn("bg-white dark:bg-black", className)}
      {...otherProps}
    />
  );
}

export function TextInput(props: TextInputProps) {
  const { className, ...otherProps } = props;

  return (
    <DefaultTextInput
      className={cn(
        "bg-white border placeholder:text-gray-600 border-gray-600 p-3 rounded-md dark:bg-black dark:text-white",
        className,
      )}
      {...otherProps}
    />
  );
}
