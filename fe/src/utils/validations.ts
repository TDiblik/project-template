import {zodResolver} from "@hookform/resolvers/zod";
import * as z from "zod";

export {zodResolver, z};

export const EmailSchema = z.email("Enter a valid email");

export const LoginFirstPageSchema = z.object({
  email: EmailSchema,
});
export type LoginFirstPageFormType = z.infer<typeof LoginFirstPageSchema>;
