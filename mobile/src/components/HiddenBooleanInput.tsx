import type React from "react";
import {useEffect} from "react";
import {Controller} from "react-hook-form";
import type {FormFieldProps} from "../utils/form";

export type HiddenBooleanInputProps = FormFieldProps & {
  value: boolean;
};

export const HiddenBooleanInput: React.FC<HiddenBooleanInputProps> = ({name, value, form}) => (
  <Controller
    name={name}
    control={form.control}
    defaultValue={value ?? false}
    render={({field: {onChange}}) => {
      // biome-ignore lint/correctness/useHookAtTopLevel: i just dont want to deal with it rn
      useEffect(() => {
        onChange(value);
      }, [onChange]);
      // biome-ignore lint/complexity/noUselessFragments: not needed
      return <></>;
    }}
  />
);
