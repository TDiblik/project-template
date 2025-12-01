import type {UseFormReturn} from "react-hook-form";

export interface FormFieldProps {
  name: string;
  form: UseFormReturn<any, any, any>;
  label?: string;
  isDisabled?: boolean;
}
