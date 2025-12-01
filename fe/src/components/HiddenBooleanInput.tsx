import {useEffect} from "react";
import {Controller, useFormContext} from "react-hook-form";
import type {FormFieldProps} from "../utils/form";

export type HiddenBooleanInputProps = FormFieldProps & {
  value: boolean;
  containerProps?: React.HTMLAttributes<HTMLDivElement>;
};

const HiddenBooleanInputInner: React.FC<{
  value: boolean;
  onChange: (value: boolean) => void;
  fieldValue: boolean;
  ref: React.Ref<HTMLDivElement>;
  rest: any;
  containerProps?: React.HTMLAttributes<HTMLDivElement>;
}> = ({value, onChange, fieldValue, ref, rest, containerProps}) => {
  useEffect(() => {
    onChange(value);
  }, [value, onChange]);

  return (
    <div className="hidden" ref={ref} {...containerProps}>
      <input type="hidden" value={fieldValue ? "true" : "false"} onChange={(e) => onChange(e.target.value === "true")} {...rest} />
    </div>
  );
};

export const HiddenBooleanInput: React.FC<HiddenBooleanInputProps> = ({name, value, containerProps}) => {
  const form = useFormContext();

  return (
    <Controller
      name={name}
      control={form.control}
      defaultValue={value ?? false}
      render={({field}) => (
        <HiddenBooleanInputInner
          value={value}
          onChange={field.onChange}
          fieldValue={field.value}
          ref={field.ref}
          rest={field}
          containerProps={containerProps}
        />
      )}
    />
  );
};
