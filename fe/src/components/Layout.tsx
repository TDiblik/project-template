import React, {useState} from "react";
import {routes} from "../utils/routes";
import {Link, matchPath, useLocation} from "react-router";
import {ThemePosibilities, useThemeStore, type ThemePosibilitiesType} from "../stores/ThemeStore";
import {useTranslation} from "react-i18next";
import {SupportedLanguages, type SupportedLanguagesType} from "../utils/i18n";
import {AnimatePresence, motion, type HTMLMotionProps} from "motion/react";

const Layout: React.FC<React.PropsWithChildren> = ({children}) => {
  const location = useLocation();
  const {i18n, t} = useTranslation();
  const {theme, changeTheme} = useThemeStore();

  const menuItems = [
    {name: t("layout.dashboard"), path: routes.index},
    {name: t("layout.profile"), path: routes.profile},
    {name: t("layout.settings"), path: routes.settings},
  ];

  const [themeOpen, setThemeOpen] = useState(false);
  const [languageOpen, setLanguageOpen] = useState(false);
  const toggleMenuClasses = "absolute left-full top-0 menu rounded-box w-48 bg-base-100 p-2 shadow";
  const toggleMenuAnimation = {
    initial: {opacity: 0, y: -8},
    animate: {opacity: 1, y: 0},
    exit: {opacity: 0, y: -8},
    transition: {duration: 0.15, ease: "easeOut"},
  } as HTMLMotionProps<"ul">;
  const toggleMenu = (menu: "language" | "theme") => {
    if (menu === "language") {
      setLanguageOpen((prev) => !prev);
      setThemeOpen(false);
    } else {
      setThemeOpen((prev) => !prev);
      setLanguageOpen(false);
    }
  };

  const changeThemeAndClose = (newTheme: ThemePosibilitiesType) => {
    changeTheme(newTheme);
    setThemeOpen(false);
  };
  const changeLanguageAndClose = (lang: SupportedLanguagesType) => {
    // todo: change into a store and sent PATCH to BE on change
    i18n.changeLanguage(lang);
    setLanguageOpen(false);
  };

  return (
    <div className="flex h-screen bg-base-200">
      {/* Sidebar */}
      <div className="flex w-64 flex-col border-r border-base-300 bg-base-100">
        <div className="border-b border-base-300 p-6 text-2xl font-bold">{t("projectName")}</div>

        {/* Menu */}
        <ul className="menu flex-1 gap-2 p-4">
          {menuItems.map((item) => (
            <li key={item.name} className={matchPath(item.path, location.pathname) ? "bg-primary text-primary-content rounded-lg" : ""}>
              <Link to={item.path}>{item.name}</Link>
            </li>
          ))}
        </ul>

        {/* Avatar & Settings */}
        <div className="border-t border-base-300 p-4">
          <div className="dropdown dropdown-top dropdown-end w-full">
            <div tabIndex={0} className="cursor-pointer flex items-center">
              <div className="btn btn-ghost btn-circle avatar">
                <div className="w-12 rounded-full">
                  <img src="https://i.pravatar.cc/300" alt={t("layout.userAvatarAlt")} />
                </div>
              </div>
              <span className="ml-2 text-base font-medium normal-case">Username (todo)</span>
            </div>

            <ul tabIndex={0} className="dropdown-content menu rounded-box z-[1] mb-2 w-52 bg-base-100 p-2 shadow">
              <li>
                <Link to={routes.settings}>{t("layout.settings")}</Link>
              </li>

              <li className="relative">
                <button className="flex justify-between w-full" onClick={() => toggleMenu("language")}>
                  {t("layout.changeLanguage.label")}
                </button>
                <AnimatePresence>
                  {languageOpen && (
                    <motion.ul {...toggleMenuAnimation} className={toggleMenuClasses}>
                      {SupportedLanguages.map((lang) => (
                        <li key={lang}>
                          <a className={i18n.language === lang ? "font-bold text-primary" : ""} onClick={() => changeLanguageAndClose(lang)}>
                            {t(`layout.changeLanguage.${lang}`)}
                          </a>
                        </li>
                      ))}
                    </motion.ul>
                  )}
                </AnimatePresence>
              </li>

              <li className="relative">
                <button className="flex justify-between w-full" onClick={() => toggleMenu("theme")}>
                  {t("layout.changeTheme.label")}
                </button>
                <AnimatePresence>
                  {themeOpen && (
                    <motion.ul {...toggleMenuAnimation} className={toggleMenuClasses}>
                      {ThemePosibilities.map((s) => (
                        <li key={s}>
                          <a className={theme === s ? "font-bold text-primary" : ""} onClick={() => changeThemeAndClose(s)}>
                            {t(`layout.changeTheme.${s}`)}
                          </a>
                        </li>
                      ))}
                    </motion.ul>
                  )}
                </AnimatePresence>
              </li>

              <div className="my-1 border-t border-base-300"></div>

              <li>
                <Link to={routes.logout} className="text-error">
                  {t("layout.logout")}
                </Link>
              </li>
            </ul>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="flex-1 overflow-auto p-6">{children}</div>
    </div>
  );
};

export default Layout;
