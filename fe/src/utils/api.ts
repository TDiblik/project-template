import {Configuration, ApiV1AuthOauthApi, ApiV1AuthOauthRedirectApi} from "@shared/api-client";
import {constants} from "./constants";

const config = new Configuration({
  basePath: constants.API_BASE_PATH,
  // optional: add default headers
  // headers: { Authorization: `Bearer ${token}` }
});

export const oAuthController = new ApiV1AuthOauthApi(config);
export const oAuthRedirectController = new ApiV1AuthOauthRedirectApi(config);

console.log(config);
