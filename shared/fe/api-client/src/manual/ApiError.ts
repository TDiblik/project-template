import {
  type GithubComTDiblikProjectTemplateApiUtilsErrorResponseType,
  GithubComTDiblikProjectTemplateApiUtilsErrorResponseTypeFromJSON,
} from "../generated";

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
          result.Body =
            GithubComTDiblikProjectTemplateApiUtilsErrorResponseTypeFromJSON(
              jsonData,
            );
        } catch {
          result.Body = { message: "Failed to parse error response" } as any;
        }
      } else {
        result.Body = { message: "No JSON body in error response" } as any;
      }
    } else {
      result.Body = { message: "No response in error object" } as any;
    }
  } catch {
    result.Body = { message: "Unknown error" } as any;
  }

  return result;
};
