import type {GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponseRedirectBackToAfterOauthEnum} from "@shared/api-client";

export const routes = {
  index: "/",
  login: "/login",
  loginOAuthRedired: "/login/oauth/redirect",
  logout: "/logout",
  profile: "/profile",
  settings: "/settings",
};

export const RedirectBackToAfterOauthToRouteMap: Record<
  GithubComTDiblikProjectTemplateApiHandlersOAuthPostReturnHandlerResponseRedirectBackToAfterOauthEnum,
  string
> = {
  index: routes.index,
  profile: routes.profile,
  settings: routes.settings,
};
