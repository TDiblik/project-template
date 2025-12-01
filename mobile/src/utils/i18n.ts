import AsyncStorage from "@react-native-async-storage/async-storage";
import * as Localization from "expo-localization";
import i18n from "i18next";
import {initReactI18next} from "react-i18next";
import translationCz from "../../assets/locales/cs/translation.json";
import translationEn from "../../assets/locales/en/translation.json";
import {constants} from "./constants";

const resources = {
  en: {translation: translationEn},
  cs: {translation: translationCz},
};

const initI18n = async () => {
  let savedLanguage = await AsyncStorage.getItem("language");
  if (!savedLanguage) {
    savedLanguage = Localization.getLocales()[0].languageCode;
  }
  if (!SupportedLanguages.some((s) => s === savedLanguage)) {
    savedLanguage = constants.DEFAULT_FALLBACK_LANGUAGE;
  }

  i18n.use(initReactI18next).init({
    resources,
    debug: constants.DEBUG,
    lng: savedLanguage,
  });
};
initI18n();

export type SupportedLanguagesType = "cs" | "en";
export const SupportedLanguages: SupportedLanguagesType[] = ["cs", "en"];

export default i18n;
