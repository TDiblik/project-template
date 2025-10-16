import "./index.css";
import "./utils/i18n.ts";
import {StrictMode} from "react";
import {createRoot} from "react-dom/client";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {BrowserRouter, Route, Routes} from "react-router";
import {routes} from "./utils/routes.ts";
import Home from "./pages/Home/Home.tsx";
import Login from "./pages/Login/Login.tsx";
import Logout from "./pages/Logout/Logout.tsx";
import OAuthRedirect from "./pages/Login/OauthRedirect.tsx";
import {LoaderProvider} from "./components/LoadingProvider.tsx";
import {ThemeProvider} from "./components/ThemeProvider.tsx";
import Settings from "./pages/Settings/Settings.tsx";
import {LoggedUserProvider} from "./components/LoggedUserProvider.tsx";

const queryClient = new QueryClient();
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <LoaderProvider>
          <ThemeProvider>
            <LoggedUserProvider>
              <Routes>
                <Route path={routes.login} element={<Login />} />
                <Route path={routes.loginOAuthRedirect} element={<OAuthRedirect />} />
                <Route path={routes.logout} element={<Logout />} />

                <Route path={routes.index} element={<Home />} />
                <Route path={routes.settings} element={<Settings />} />
              </Routes>
            </LoggedUserProvider>
          </ThemeProvider>
        </LoaderProvider>
      </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>,
);
