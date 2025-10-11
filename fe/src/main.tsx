import "./index.css";
import "./utils/i18n.ts";
import {StrictMode} from "react";
import {createRoot} from "react-dom/client";
import {BrowserRouter, Route, Routes} from "react-router";
import {routes} from "./utils/routes.ts";
import Home from "./pages/Home/Home.tsx";
import Login from "./pages/Login/Login.tsx";
import Logout from "./pages/Logout/Logout.tsx";
import OAuthRedirect from "./pages/Login/OauthRedirect.tsx";
import {LoaderProvider} from "./components/LoadingProvider.tsx";
import {ThemeProvider} from "./components/ThemeProvider.tsx";
import Settings from "./pages/Settings/Settings.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <LoaderProvider>
        <ThemeProvider>
          <Routes>
            <Route path={routes.login} element={<Login />} />
            <Route path={routes.loginOAuthRedired} element={<OAuthRedirect />} />
            <Route path={routes.logout} element={<Logout />} />

            <Route path={routes.index} element={<Home />} />
            <Route path={routes.settings} element={<Settings />} />
          </Routes>
        </ThemeProvider>
      </LoaderProvider>
    </BrowserRouter>
  </StrictMode>,
);
