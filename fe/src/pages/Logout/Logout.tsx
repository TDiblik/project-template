import React from "react";
import {useTranslation} from "react-i18next";
import {constants} from "../../utils/constants";
import {routes} from "../../utils/routes";

const Logout: React.FC = () => {
  const {t} = useTranslation();

  React.useEffect(() => {
    // easiest way to reset every state
    localStorage.removeItem(constants.LOCAL_STORAGE_TOKEN_KEY);
    window.location.href = routes.login;
  }, []);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-base-100">
      <span className="loading loading-spinner loading-xl text-primary mb-4"></span>
      <p className="text-lg font-medium text-base-content">{t("logoutPage.title")}</p>
      <p className="text-sm text-base-content/70 mt-1">{t("logoutPage.description")}</p>
    </div>
  );
};
export default Logout;
