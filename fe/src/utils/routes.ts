import type {RedirectBackToAfterOauthEnum} from "@shared/api-client";

export const routes = {
  index: "/",
  login: "/login",
  loginOAuthRedirect: "/login/oauth/redirect",
  logout: "/logout",
  settings: "/settings",
};

export const RedirectBackToAfterOauthToRouteMap: Record<RedirectBackToAfterOauthEnum, string> = {
  index: routes.index,
  settings: routes.settings,
};
