import {
  Configuration,
  ApiV1AuthOauthApi,
  ApiV1AuthOauthRedirectApi,
  ApiV1AuthApi,
  type GithubComTDiblikProjectTemplateApiUtilsErrorResponseType,
  GithubComTDiblikProjectTemplateApiUtilsErrorResponseTypeFromJSON,
} from "@shared/api-client";
import {constants} from "./constants";

const config = new Configuration({
  basePath: constants.API_BASE_PATH,
  // optional: add default headers
  // headers: { Authorization: `Bearer ${token}` }
});

export const AuthController = new ApiV1AuthApi(config);
export const oAuthController = new ApiV1AuthOauthApi(config);
export const oAuthRedirectController = new ApiV1AuthOauthRedirectApi(config);

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
