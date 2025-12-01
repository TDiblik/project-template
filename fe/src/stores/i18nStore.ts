import type {TranslationPosibilitiesType} from "@shared/api-client";
import i18n, {type i18n as i18nType} from "i18next";
import {create} from "zustand";
import {UserController} from "../utils/api";
import {useAuthTokenStore} from "./TokenStore";

interface i18nStoreState {
  i18n: i18nType;
  changeLanguage: (newLanguage: TranslationPosibilitiesType) => void;
}
export const usei18nStore = create<i18nStoreState>()(() => ({
  i18n: i18n,
  changeLanguage: (newLanguage) => {
    i18n.changeLanguage(newLanguage);
    if (useAuthTokenStore.getState().isAuthenticatedAndLoaded()) {
      UserController.apiV1PrivateUserMePatch({
        githubComTDiblikProjectTemplateApiHandlersPatchUserMeHandlerRequest: {
          preferedLanguage: newLanguage,
        },
      });
    }
  },
}));
