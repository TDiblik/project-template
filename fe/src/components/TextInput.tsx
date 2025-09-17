import {Controller, useFormContext} from "react-hook-form";
import type {FormFieldProps} from "../utils/form";

export type TextInputProps = FormFieldProps & {
  placeholder?: string;
};

export const TextInput: React.FC<TextInputProps> = (props) => {
  const form = useFormContext();

  return (
    <Controller
      name={props.name}
      control={form.control}
      render={({field, fieldState}) => {
        const {ref, ...rest} = field;
        const hasError = !!fieldState.error;

        return (
          <div className="form-control w-full" ref={ref}>
            {props.label && (
              <label className="label mb-1">
                <span className={`label-text ${hasError ? "text-red-500" : ""}`}>{props.label}</span>
              </label>
            )}
            <input
              type="text"
              placeholder={props.placeholder}
              disabled={props.isDisabled}
              className={`input input-bordered w-full ${hasError ? "input-error" : ""}`}
              {...rest}
            />
            {hasError && <span className="text-red-500 text-sm mt-1">{fieldState.error?.message}</span>}
          </div>
        );
      }}
    />
  );
};
