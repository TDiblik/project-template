import {t} from "i18next";
import type {ApiError} from "@shared/api-client";

export const TranslateApiErrorMessage = (error: ApiError) => t(error.Body.msg || "be.error.internal_server_error");
