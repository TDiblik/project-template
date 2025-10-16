import {create} from "zustand";
import {getRawJWT, type IAuthToken, parseJWT} from "../utils/token";
import {constants} from "../utils/constants";
import {useLoggedUserStore} from "./LoggedUserStore";

interface TokenStoreState {
  tokenRaw: string | null;
  token: () => IAuthToken | null;
  setToken: (newToken: string) => void;
  resetToken: () => void;
  isAuthenticated: () => boolean;
  isAuthenticatedAndLoaded: () => boolean;
}

export const useAuthTokenStore = create<TokenStoreState>()((set, get) => ({
  tokenRaw: getRawJWT(),
  token: () => parseJWT(get().tokenRaw),
  setToken: (newToken) => {
    localStorage.setItem(constants.LOCAL_STORAGE_TOKEN_KEY, newToken);
    set(() => ({tokenRaw: newToken}));
  },
  resetToken: () => {
    localStorage.removeItem(constants.LOCAL_STORAGE_TOKEN_KEY);
    set(() => ({tokenRaw: undefined}));
  },
  isAuthenticated: () => {
    const tokenParsed = get().token();
    if (!tokenParsed) return false;
    return new Date(tokenParsed.exp * 1000) > new Date();
  },
  isAuthenticatedAndLoaded: () => get().isAuthenticated() && !!useLoggedUserStore.getState().user,
}));
