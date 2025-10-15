import React, {useEffect} from "react";
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
      // eslint-disable-next-line react-hooks/rules-of-hooks
      useEffect(() => {
        onChange(value);
      }, [onChange]);
      return <></>;
    }}
  />
);
