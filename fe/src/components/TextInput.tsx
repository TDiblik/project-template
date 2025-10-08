import {Controller, useFormContext} from "react-hook-form";
import type {FormFieldProps} from "../utils/form";

export type TextInputProps = FormFieldProps & {
  placeholder?: string;
  containerProps?: React.HTMLAttributes<HTMLDivElement>;
  inputProps?: React.InputHTMLAttributes<HTMLInputElement>;
  labelProps?: React.LabelHTMLAttributes<HTMLLabelElement>;
  labelSpanProps?: React.HTMLAttributes<HTMLSpanElement>;
  labelSpanAdditionalClassname?: string;
  errorSpanProps?: React.HTMLAttributes<HTMLSpanElement>;
  errorSpanPropsAdditionalClassname?: string;
  isOptional?: boolean;
  hasBigText?: boolean;
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
          <div className="form-control w-full" ref={ref} {...props.containerProps}>
            <fieldset className="fieldset">
              <legend className="fieldset-legend pb-1">
                {props.label && (
                  <label className={`${props.hasBigText && "text-base font-normal"} label`} {...props.labelProps}>
                    <span className={`label-text ${hasError ? "text-red-500" : ""} ${props.labelSpanAdditionalClassname}`} {...props.labelSpanProps}>
                      {props.label}
                    </span>
                  </label>
                )}
              </legend>
              <input
                type="text"
                placeholder={props.placeholder}
                disabled={props.isDisabled}
                className={`input input-bordered w-full ${hasError ? "input-error" : ""}`}
                {...props.inputProps}
                {...rest}
              />
              {props.isOptional && <p className="label">Optional</p>}
            </fieldset>

            {hasError && (
              <span className={`text-red-500 text-sm mt-1 ${props.errorSpanPropsAdditionalClassname}`} {...props.errorSpanProps}>
                {fieldState.error?.message}
              </span>
            )}
          </div>
        );
      }}
    />
  );
};
