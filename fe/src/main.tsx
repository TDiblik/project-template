import "./index.css";
import {StrictMode} from "react";
import {createRoot} from "react-dom/client";
import {BrowserRouter, Route, Routes} from "react-router";
import {routes} from "./utils/routes.ts";
import Home from "./pages/Home/Home.tsx";
import Login from "./pages/Login/Login.tsx";
import Logout from "./pages/Logout/Logout.tsx";
import OAuthRedirect from "./pages/Login/OauthRedirect.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path={routes.index} element={<Home />} />
        <Route path={routes.login} element={<Login />} />
        <Route path={routes.loginOAuthRedired} element={<OAuthRedirect />} />
        <Route path={routes.logout} element={<Logout />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
);
