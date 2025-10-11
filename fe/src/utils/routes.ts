import type {RedirectBackToAfterOauthEnum} from "./api";

export const routes = {
  index: "/",
  login: "/login",
  loginOAuthRedired: "/login/oauth/redirect",
  logout: "/logout",
  profile: "/profile",
  settings: "/settings",
};

export const RedirectBackToAfterOauthToRouteMap: Record<RedirectBackToAfterOauthEnum, string> = {
  index: routes.index,
  profile: routes.profile,
  settings: routes.settings,
};
