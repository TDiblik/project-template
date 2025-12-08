import type {ThemePosibilitiesType} from "@shared/api-client";
import {create} from "zustand";
import {UserController} from "../utils/api";
import {constants} from "../utils/constants";
import {useAuthTokenStore} from "./TokenStore";

interface ThemeStoreState {
  theme: ThemePosibilitiesType;
  changeTheme: (theme: ThemePosibilitiesType) => void;
}

export const useThemeStore = create<ThemeStoreState>()((set) => ({
  theme: (localStorage.getItem(constants.LOCAL_STORAGE_THEME_KEY) as ThemePosibilitiesType) ?? document.documentElement.getAttribute("data-theme") ?? "light",
  changeTheme: (theme) => {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem(constants.LOCAL_STORAGE_THEME_KEY, theme);
    set(() => ({theme}));
    if (useAuthTokenStore.getState().isAuthenticatedAndLoaded()) {
      UserController.apiV1PrivateUserMePatch({
        githubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest: {
          preferedTheme: theme,
        },
      });
    }
  },
}));
