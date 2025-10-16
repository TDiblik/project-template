import i18n from "i18next";
import {initReactI18next} from "react-i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import {constants} from "./constants";
import Backend from "i18next-http-backend";

i18n
  .use(Backend)
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    load: "languageOnly",
    debug: constants.DEBUG,
    lng: (localStorage.getItem(constants.LOCAL_STORAGE_LOCALIZATION_KEY) ?? "").split("-")[0],
    fallbackLng: constants.DEFAULT_FALLBACK_LANGUAGE,
    detection: {
      lookupLocalStorage: constants.LOCAL_STORAGE_LOCALIZATION_KEY,
    },
  });

export default i18n;
