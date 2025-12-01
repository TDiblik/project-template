import type {ApiError} from "@shared/api-client";
import {t} from "i18next";

export const TranslateApiErrorMessage = (error: ApiError) => t(error.Body.msg || "be.error.internal_server_error");
