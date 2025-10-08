import {useEffect, useState} from "react";
import {useNavigate, useSearchParams} from "react-router";
import {oAuthController} from "../../utils/api";
import {useAuthTokenStore} from "../../stores/TokenStore";
import {RedirectBackToAfterOauthToRouteMap, routes} from "../../utils/routes";
import {useTranslation} from "react-i18next";
import {AnimatePresence, motion, type HTMLMotionProps} from "motion/react";

const delayed = (delay: number) =>
  ({
    initial: {opacity: 0, y: 5},
    animate: {opacity: 1, y: 0},
    transition: {delay},
  }) as HTMLMotionProps<"p"> | HTMLMotionProps<"h1">;

const OAuthRedirect = () => {
  const {t} = useTranslation();
  const navigate = useNavigate();
  const [query] = useSearchParams();
  const oAuthCode = query.get("code");
  const oAuthState = query.get("state");
  const [shouldFail, setShouldFail] = useState(!oAuthCode || !oAuthState);

  const {setToken} = useAuthTokenStore();

  useEffect(() => {
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
        navigate(RedirectBackToAfterOauthToRouteMap[s.redirectBackToAfterOauth]);
      })
      .catch((error) => {
        console.log(error); // todo: add better error logging
        setShouldFail(true);
      });
  }, [oAuthCode, oAuthState]);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-base-100 text-center px-4">
      <AnimatePresence mode="wait">
        {!shouldFail ? (
          <motion.div
            key="loading"
            initial={{opacity: 0, y: 10}}
            animate={{opacity: 1, y: 0}}
            exit={{opacity: 0, y: -10}}
            transition={{duration: 0.6, ease: "easeInOut"}}
            className="flex flex-col items-center"
          >
            <motion.span
              className="loading loading-spinner loading-xl text-primary mb-4"
              animate={{
                scale: [1, 1.15, 1],
                opacity: [1, 0.7, 1],
              }}
              transition={{
                duration: 1.2,
                repeat: Infinity,
                ease: "easeInOut",
              }}
            />
            <motion.p className="text-lg font-medium text-base-content" {...delayed(0.2)}>
              {t("oauthRedirectPage.loginInProgress.title")}
            </motion.p>
            <motion.p className="text-sm text-base-content/70 mt-1" {...delayed(0.4)}>
              {t("oauthRedirectPage.loginInProgress.description")}
            </motion.p>
          </motion.div>
        ) : (
          <motion.div
            key="fail"
            initial={{opacity: 0}}
            animate={{opacity: 1}}
            exit={{opacity: 0}}
            transition={{duration: 0.8, ease: "easeOut"}}
            className="flex flex-col items-center"
          >
            <motion.h1 className="text-2xl font-bold text-error mb-2" {...delayed(0.2)}>
              {t("oauthRedirectPage.loginFailed.title")}
            </motion.h1>
            <motion.p className="text-base text-base-content/80 mb-6" {...delayed(0.4)}>
              {t("oauthRedirectPage.loginFailed.description")}
            </motion.p>
            <motion.button whileHover={{scale: 1.05}} whileTap={{scale: 0.95}} className="btn btn-primary" onClick={() => navigate(routes.login)}>
              {t("oauthRedirectPage.loginFailed.button")}
            </motion.button>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
};

export default OAuthRedirect;
