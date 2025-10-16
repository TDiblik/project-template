import type {RedirectBackToAfterOauthEnum} from "@shared/api-client";

export const routes = {
  index: "/",
  login: "/login",
  loginOAuthRedirect: "/login/oauth/redirect",
  logout: "/logout",
  profile: "/profile",
  settings: "/settings",
};

export const RedirectBackToAfterOauthToRouteMap: Record<RedirectBackToAfterOauthEnum, string> = {
  index: routes.index,
  profile: routes.profile,
  settings: routes.settings,
};
