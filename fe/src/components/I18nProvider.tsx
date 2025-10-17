import {useEffect} from "react";
import {useLoggedUser} from "../stores/LoggedUserStore";
import type {TranslationPosibilitiesType} from "@shared/api-client";
import {useAuthTokenStore} from "../stores/TokenStore";
import {usei18nStore} from "../stores/i18nStore";
import {useTranslation} from "react-i18next";

export const I18nProvider: React.FC<React.PropsWithChildren> = ({children}) => {
  const {loggedUser} = useLoggedUser();
  const {changeLanguage} = usei18nStore();
  const {isAuthenticatedAndLoaded} = useAuthTokenStore();
  const {i18n} = useTranslation();

  const _language = i18n.language as TranslationPosibilitiesType;
  useEffect(() => {
    if (i18n.language !== _language) {
      changeLanguage(_language);
    }
  }, [_language]);
  useEffect(() => {
    if (!isAuthenticatedAndLoaded()) return;
    if (!loggedUser?.preferedLanguage) {
      changeLanguage(_language as TranslationPosibilitiesType);
    } else if (loggedUser?.preferedLanguage && loggedUser?.preferedLanguage !== _language) {
      changeLanguage(loggedUser?.preferedLanguage as TranslationPosibilitiesType);
    }
  }, [loggedUser]);

  return children;
};
