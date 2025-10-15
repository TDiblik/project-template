import i18n, {type i18n as i18nType} from "i18next";
import {create} from "zustand";
import type {SupportedLanguagesType} from "../utils/i18n";

interface i18nStoreState {
  i18n: i18nType;
  language: () => SupportedLanguagesType;
  changeLanguage: (newLanguage: SupportedLanguagesType) => void;
}
export const usei18nStore = create<i18nStoreState>()(() => ({
  i18n: i18n,
  language: () => i18n.language as SupportedLanguagesType,
  changeLanguage: (newLanguage) => {
    i18n.changeLanguage(newLanguage);
    // todo: send PATCH to BE on change
  },
}));
