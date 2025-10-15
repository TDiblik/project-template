import {type ApiError, Configuration, getAuthController, getoAuthController, getoAuthRedirectController, getUserController} from "@shared/api-client";
import {constants} from "./constants";
import {t} from "i18next";

const config = new Configuration({
  basePath: constants.API_BASE_PATH,
  // apiKey: () => useAuthTokenStore.getState().tokenRaw ?? "",
});

export const AuthController = getAuthController(config);
export const oAuthController = getoAuthController(config);
export const oAuthRedirectController = getoAuthRedirectController(config);
export const UserController = getUserController(config);

export const TranslateApiErrorMessage = (error: ApiError) => t(error.Body.msg || "be.error.internal_server_error");
