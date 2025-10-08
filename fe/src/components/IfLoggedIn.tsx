import {useNavigate} from "react-router";
import {useAuthTokenStore} from "../stores/TokenStore";
import React from "react";
import {routes} from "../utils/routes";

interface IfLoggedInProps {
  redirectToLogin?: boolean;
  children?: React.ReactNode;
}
export const IfLoggedIn: React.FC<IfLoggedInProps> = (props) => {
  const navigate = useNavigate();
  const {token, tokenRaw} = useAuthTokenStore();
  const [canView, setCanView] = React.useState(true);

  React.useEffect(() => {
    if (!tokenRaw || !token()) {
      setCanView(false);
      if (props.redirectToLogin) {
        navigate(routes.login);
      }
    }
  }, [token, tokenRaw]);

  return <>{canView ? props.children : null}</>;
};
