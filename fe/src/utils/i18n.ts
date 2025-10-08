import i18n from "i18next";
import {initReactI18next} from "react-i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import {constants} from "./constants";
import Backend from "i18next-http-backend";

i18n.use(Backend).use(LanguageDetector).use(initReactI18next).init({
  debug: constants.DEBUG,
  fallbackLng: constants.DEFAULT_FALLBACK_LANGUAGE,
});

export default i18n;
