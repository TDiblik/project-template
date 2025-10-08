const apiBasePathRaw = import.meta.env.VITE_API_BASE_PATH as string;
export const constants = {
  DEBUG: import.meta.env.NODE_ENV !== "production",

  API_BASE_PATH: apiBasePathRaw.endsWith("/") ? apiBasePathRaw.slice(0, -1) : apiBasePathRaw,
  GIT_TAG: import.meta.env.VITE_GIT_TAG as string,

  DEFAULT_FALLBACK_LANGUAGE: import.meta.env.VITE_DEFAULT_FALLBACK_LANGUAGE as string,

  TOKEN_HEADER_KEY: "x-user-token",
  LOCAL_STORAGE_TOKEN_KEY: "project-template-auth-token",
  LOCAL_STORAGE_THEME_KEY: "project-template-theme",
  LOCAL_STORAGE_LOCALIZATION_KEY: "project-template-localization",
};
