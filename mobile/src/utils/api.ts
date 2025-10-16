import {Configuration, getAuthController, getoAuthController, getoAuthRedirectController, getUserController} from "@shared/api-client";
import {constants} from "./constants";

const config = new Configuration({
  basePath: constants.API_BASE_PATH,
  // apiKey: () => useAuthTokenStore.getState().tokenRaw ?? "",
});

export const AuthController = getAuthController(config);
export const oAuthController = getoAuthController(config);
export const oAuthRedirectController = getoAuthRedirectController(config);
export const UserController = getUserController(config);
