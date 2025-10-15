import {Redirect, SplashScreen} from "expo-router";
import {useEffect} from "react";

SplashScreen.preventAutoHideAsync();
export default function HomeScreen() {
  //   let [fontsLoaded] = useFonts({Raleway_400Regular, Raleway_500Medium});
  //   React.useEffect(() => {
  //     if (fontsLoaded) {
  //       SplashScreen.hideAsync();
  //     }
  //   }, [fontsLoaded]);
  useEffect(() => {
    SplashScreen.hideAsync();
  }, []);

  return <Redirect href="/login" />;
}
