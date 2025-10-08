import {create} from "zustand";
import {constants} from "../utils/constants";

export type ThemePosibilitiesType = "light" | "dark";
export const ThemePosibilities: ThemePosibilitiesType[] = ["light", "dark"];
interface ThemeStoreState {
  theme: ThemePosibilitiesType;
  changeTheme: (theme: ThemePosibilitiesType) => void;
}

export const useThemeStore = create<ThemeStoreState>()((set) => ({
  theme:
    (localStorage.getItem(constants.LOCAL_STORAGE_THEME_KEY) as ThemePosibilitiesType) ??
    document.documentElement.getAttribute("data-theme") ??
    "light",
  changeTheme: (theme) => {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem(constants.LOCAL_STORAGE_THEME_KEY, theme);
    set(() => ({theme}));
  },
}));
