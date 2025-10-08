import {useEffect} from "react";
import type {UseFormReturn, FieldValues} from "react-hook-form";
import {constants} from "./constants";

export function useFormLog<T extends FieldValues>(form: UseFormReturn<T>) {
  useEffect(() => {
    if (constants.DEBUG) {
      const subscription = form.watch(() => {
        console.log("Form Values:", form.getValues());
        console.log("Form Errors:", form.formState.errors);
      });
      return () => subscription.unsubscribe();
    }
  }, [form]);
}
