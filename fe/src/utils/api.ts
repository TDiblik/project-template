import {
  ApiV1PrivateUserApi,
  ApiV1PublicAuthApi,
  ApiV1PublicAuthOauthApi,
  ApiV1PublicAuthOauthRedirectApi,
  Configuration,
  GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponseRedirectBackToAfterOauthEnum,
  type GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse,
  type GithubComTDiblikProjectTemplateApiUtilsErrorResponseType,
  GithubComTDiblikProjectTemplateApiUtilsErrorResponseTypeFromJSON,
} from "@shared/api-client";
import {constants} from "./constants";
import {t} from "i18next";
import {useAuthTokenStore} from "../stores/TokenStore";

const config = new Configuration({
  basePath: constants.API_BASE_PATH,
  apiKey: () => useAuthTokenStore.getState().tokenRaw ?? "",
});

export const AuthController = new ApiV1PublicAuthApi(config);
export const oAuthController = new ApiV1PublicAuthOauthApi(config);
export const oAuthRedirectController = new ApiV1PublicAuthOauthRedirectApi(config);
export const UserController = new ApiV1PrivateUserApi(config);

export interface ApiError {
  Ok: boolean;
  StatusCode: number;
  Body: GithubComTDiblikProjectTemplateApiUtilsErrorResponseType;
}
export const ConvertToApiError = async (error: any): Promise<ApiError> => {
  const result: ApiError = {
    Ok: false,
    StatusCode: 500,
    Body: {
      msg: "",
      status: "",
    },
  };

  try {
    if (error?.response) {
      if (typeof error.response.ok === "boolean") {
        result.Ok = error.response.ok;
      }

      if (typeof error.response.status === "number") {
        result.StatusCode = error.response.status;
      }

      if (typeof error.response.json === "function") {
        try {
          const jsonData = await error.response.json();
          result.Body = GithubComTDiblikProjectTemplateApiUtilsErrorResponseTypeFromJSON(jsonData);
        } catch {
          result.Body = {message: "Failed to parse error response"} as any;
        }
      } else {
        result.Body = {message: "No JSON body in error response"} as any;
      }
    } else {
      result.Body = {message: "No response in error object"} as any;
    }
  } catch {
    result.Body = {message: "Unknown error"} as any;
  }

  return result;
};
export const TranslateApiErrorMessage = (error: ApiError) => t(error.Body.msg || "be.error.internal_server_error");

export type OauthRedirectHandlerResponse = GithubComTDiblikProjectTemplateApiHandlersOauthRedirectHandlerResponse;
export type OauthRedirectHandlerRequest = Promise<OauthRedirectHandlerResponse>;
export type RedirectBackToAfterOauthEnum = GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponseRedirectBackToAfterOauthEnum;
