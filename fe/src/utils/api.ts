import {Configuration, getAuthController, getoAuthController, getoAuthRedirectController, getUserController} from "@shared/api-client";
import {constants} from "./constants";
import {useAuthTokenStore} from "../stores/TokenStore";

const config = new Configuration({
  basePath: constants.API_BASE_PATH,
  apiKey: () => useAuthTokenStore.getState().tokenRaw ?? "",
  // todo: add a middleware that refreshes the api token
});

export const AuthController = getAuthController(config);
export const oAuthController = getoAuthController(config);
export const oAuthRedirectController = getoAuthRedirectController(config);
export const UserController = getUserController(config);
