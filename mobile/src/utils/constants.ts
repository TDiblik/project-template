import {ApiConstants} from "@shared/api-client";

const apiBasePathRaw = process.env.EXPO_PUBLIC_API_BASE_PATH as string;
export const constants = {
  DEBUG: process.env.EXPO_PUBLIC_ENV !== "production",

  API_BASE_PATH: apiBasePathRaw.endsWith("/") ? apiBasePathRaw.slice(0, -1) : apiBasePathRaw,
  DEFAULT_FALLBACK_LANGUAGE: process.env.EXPO_PUBLIC_DEFAULT_FALLBACK_LANGUAGE as string,

  ...ApiConstants,
};
