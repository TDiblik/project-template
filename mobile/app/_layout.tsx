import "../src/global.css";
import "../src/utils/i18n";
import {Slot} from "expo-router";
import {StrictMode} from "react";

export default function Layout() {
  return (
    <StrictMode>
      <Slot />
    </StrictMode>
  );
}
