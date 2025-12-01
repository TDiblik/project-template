import type {ThemePosibilitiesType} from "@shared/api-client";
import {useEffect} from "react";
import {useLoggedUser} from "../stores/LoggedUserStore";
import {useThemeStore} from "../stores/ThemeStore";
import {useAuthTokenStore} from "../stores/TokenStore";

export const ThemeProvider: React.FC<React.PropsWithChildren> = ({children}) => {
  const {isAuthenticatedAndLoaded} = useAuthTokenStore();
  const {loggedUser} = useLoggedUser();
  const {theme, changeTheme} = useThemeStore();
  useEffect(() => {
    if (document.documentElement.getAttribute("data-theme") !== theme) {
      changeTheme(theme);
    }
  }, []);
  useEffect(() => {
    if (!isAuthenticatedAndLoaded()) return;
    if (!loggedUser?.preferedTheme) {
      changeTheme(theme);
    } else if (loggedUser?.preferedTheme && loggedUser?.preferedTheme !== theme) {
      changeTheme(loggedUser?.preferedTheme as ThemePosibilitiesType);
    }
  }, [loggedUser]);

  return children;
};
