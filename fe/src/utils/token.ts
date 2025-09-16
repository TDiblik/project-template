import {constants} from "./constants";

export interface IAuthToken {
  sub: string;
  jti: string;
  user_id: string;
  user_email: string;
  user_first_name?: string;
  user_last_name?: string;
  user_handle?: string;
  exp: number;
  iss: string;
  aud: string;
}

export const getRawJWT = () => localStorage.getItem(constants.LOCAL_STORAGE_TOKEN_KEY);
export const parseJWT = (token: string | null): IAuthToken | null => {
  if (!token) {
    return null;
  }

  try {
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      window
        .atob(base64)
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join(""),
    );

    return JSON.parse(jsonPayload) as IAuthToken;
  } catch {
    return null;
  }
};
