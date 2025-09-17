import React from "react";
import {useNavigate, useSearchParams} from "react-router";
import {oAuthController} from "../../utils/api";
import {useAuthTokenStore} from "../../stores/TokenStore";
import {routes} from "../../utils/routes";

const OAuthRedirect = () => {
  const navigate = useNavigate();
  const [query] = useSearchParams();
  const oAuthCode = query.get("code");
  const oAuthState = query.get("state");
  const shouldFail = !oAuthCode || !oAuthState;

  const {setToken} = useAuthTokenStore();

  React.useEffect(() => {
    if (!oAuthCode || !oAuthState) {
      return;
    }
    oAuthController
      .apiV1AuthOauthReturnGet({
        state: oAuthState,
        code: oAuthCode,
      })
      .then((s) => {
        setToken(s.authToken);
        // when adding redirect url after oauth, add it here
        switch (s.redirectBackToAfterOauth) {
          case "index":
            navigate(routes.index);
            break;
          case "profile":
            navigate(routes.profile);
            break;
          case "settings":
            navigate(routes.settings);
            break;
        }
      })
      .catch(() => {});
  }, [oAuthCode, oAuthState]);

  if (shouldFail) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-base-100 px-4 text-center">
        <h1 className="text-2xl font-bold text-error mb-2">Login Failed</h1>
        <p className="text-base text-base-content/80 mb-6">We couldn't complete the sign-in process. Please try again.</p>
        <button className="btn btn-primary" onClick={() => navigate("/login")}>
          Back to Login
        </button>
      </div>
    );
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-base-100">
      <span className="loading loading-spinner loading-xl text-primary mb-4"></span>
      <p className="text-lg font-medium text-base-content">Signing you in...</p>
      <p className="text-sm text-base-content/70 mt-1">Please wait while we complete the process.</p>
    </div>
  );
};

export default OAuthRedirect;
