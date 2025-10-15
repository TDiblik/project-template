import React, {useEffect, useRef} from "react";
import {Text, TextInput as RNTextInput, View, Animated} from "react-native";
import {Controller} from "react-hook-form";
import type {FormFieldProps} from "../utils/form";

export type TextInputProps = FormFieldProps & {
  placeholder?: string;
  label?: string;
  isOptional?: boolean;
  hasBigText?: boolean;
  inputProps?: React.ComponentProps<typeof RNTextInput>;
  labelTextClassName?: string;
  errorTextClassName?: string;
};

export const TextInput: React.FC<TextInputProps> = ({
  name,
  placeholder,
  label,
  isOptional,
  hasBigText,
  inputProps,
  labelTextClassName,
  errorTextClassName,
  form,
}) => {
  const fadeAnim = useRef(new Animated.Value(0)).current;

  const hasError = !!form.formState.errors[name];
  useEffect(() => {
    Animated.timing(fadeAnim, {
      toValue: hasError ? 1 : 0,
      duration: 200,
      useNativeDriver: true,
    }).start();
  }, [fadeAnim, hasError]);

  return (
    <Controller
      name={name}
      control={form.control}
      render={({field: {onChange, onBlur, value}, fieldState}) => (
        <View className="w-full mb-4">
          {label && (
            <Text
              className={`${hasBigText ? "text-base font-normal" : "text-sm font-medium"} ${labelTextClassName} ${hasError ? "text-red-500" : "text-gray-800"}`}
            >
              {label}
              {isOptional && <Text className="text-gray-400"> (Optional)</Text>}
            </Text>
          )}

          <RNTextInput
            placeholder={placeholder}
            value={value}
            onBlur={onBlur}
            onChangeText={onChange}
            className={`border rounded-md p-3 w-full ${hasError ? "border-red-500" : "border-gray-300"}`}
            {...inputProps}
          />

          {fieldState.error?.message && (
            <Animated.Text style={{opacity: fadeAnim}}>
              <Text className={`text-red-500 text-sm mt-1 ${errorTextClassName}`}>{fieldState.error.message}</Text>
            </Animated.Text>
          )}
        </View>
      )}
    />
  );
};
