const apiBasePathRaw = import.meta.env.VITE_API_BASE_PATH as string;
export const constants = {
  DEBUG: import.meta.env.NODE_ENV !== "production",

  API_BASE_PATH: apiBasePathRaw.endsWith("/") ? apiBasePathRaw.slice(0, -1) : apiBasePathRaw,
  GIT_TAG: import.meta.env.VITE_GIT_TAG as string,
  TOKEN_HEADER_KEY: "x-user-token",
};
