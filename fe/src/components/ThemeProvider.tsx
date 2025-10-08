import {useEffect} from "react";
import {useThemeStore} from "../stores/ThemeStore";

export const ThemeProvider: React.FC<React.PropsWithChildren> = ({children}) => {
  const {theme, changeTheme} = useThemeStore();
  useEffect(() => {
    if (document.documentElement.getAttribute("data-theme") !== theme) {
      changeTheme(theme);
    }
  }, []);

  return children;
};
