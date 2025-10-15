import {
  ApiV1PrivateUserApi,
  ApiV1PublicAuthApi,
  ApiV1PublicAuthOauthApi,
  ApiV1PublicAuthOauthRedirectApi,
  Configuration,
} from "../generated";

export const getAuthController = (config: Configuration) => new ApiV1PublicAuthApi(config);
export const getoAuthController = (config: Configuration) => new ApiV1PublicAuthOauthApi(config);
export const getoAuthRedirectController = (config: Configuration) => new ApiV1PublicAuthOauthRedirectApi(config);
export const getUserController = (config: Configuration) => new ApiV1PrivateUserApi(config);
