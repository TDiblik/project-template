import {zodResolver} from "@hookform/resolvers/zod";
import * as z from "zod";

export {zodResolver, z};

// ----------- General -----------
export const EmailSchema = z.email("Enter a valid email").min(1, "Email is required");
export const PasswordSchema = z
  .string("You must enter a password")
  .min(6, "Password must be at least 6 characters long")
  .max(32, "Password cannot be longer than 32 characters")
  .regex(/[A-Z]/, "Password must contain at least one uppercase letter")
  .regex(/[a-z]/, "Password must contain at least one lowercase letter")
  .regex(/[0-9]/, "Password must contain at least one number")
  .refine((val) => new TextEncoder().encode(val).length <= 72, "Password is too long");
export const ConfirmPasswordSchema = z.string("Re-enter your password to make sure it's correct");
// -----------------------------

// ----------- Pages -----------
export const LoginFirstPageSchema = z.object({
  email: EmailSchema,
  password: PasswordSchema,
});
export type LoginPageFormType = z.infer<typeof LoginFirstPageSchema>;

export const SignUpFirstPageSchema = z
  .object({
    email: EmailSchema,
    password: PasswordSchema,
    confirmPassword: ConfirmPasswordSchema,
    useUsername: z.boolean(),
    firstName: z.string().optional(),
    lastName: z.string().optional(),
    username: z.string().optional(),
  })
  .refine(({useUsername, username}) => (useUsername ? (username?.trim().length ?? 0) > 0 : true), {path: ["username"], error: "Username is required"})
  .refine(({useUsername, username}) => (useUsername ? (username?.trim().length ?? 0) >= 3 : true), {
    path: ["username"],
    error: "Username must have 3 or more characters",
  })
  .refine(({useUsername, firstName}) => (!useUsername ? (firstName?.trim().length ?? 0) > 0 : true), {
    path: ["firstName"],
    error: "First name is required",
  })
  .refine(({useUsername, lastName}) => (!useUsername ? (lastName?.trim().length ?? 0) > 0 : true), {
    path: ["lastName"],
    error: "Last name is required",
  })
  .refine(({password, confirmPassword}) => password === confirmPassword, {path: ["confirmPassword"], error: "Passwords do not match"});
export type SignUpPageFormType = z.infer<typeof SignUpFirstPageSchema>;

export type LoginOrSignUpPageFormType = LoginPageFormType | SignUpPageFormType;
