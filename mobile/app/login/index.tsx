import FontAwesome6 from "@expo/vector-icons/FontAwesome6";
import React, {useState} from "react";
import {View, Text, TouchableOpacity, ScrollView, KeyboardAvoidingView, Platform} from "react-native";
import {useForm, FormProvider} from "react-hook-form";
import {useTranslation} from "react-i18next";
// import {useNavigation} from "@react-navigation/native";
import {
  LoginFirstPageSchema,
  SignUpFirstPageSchema,
  SignUpPageFormType,
  zodResolver,
  type LoginOrSignUpPageFormType,
} from "../../src/utils/validations";
import {AuthController, oAuthRedirectController} from "../../src/utils/api";
import {TextInput} from "../../src/components/TextInput";

// todo:
// - add register
// - implement oauth
// - implement redirection to home page
// - some forced after-login flow
// - implement auth guard
// - implement loading store and loading provider
// - add different themes

const LoginScreen = () => {
  const {t} = useTranslation();
  // const navigation = useNavigation();
  // const { setToken } = useAuthTokenStore();
  // const { setLoading } = useLoadingStore();
  const [isSignUp, setIsSignUp] = useState<boolean>(false);
  const [beErrorMessage, setBEErrorMessage] = useState<string | undefined>(undefined);

  const form = useForm<LoginOrSignUpPageFormType>({
    mode: "onChange",
    resolver: zodResolver(!isSignUp ? LoginFirstPageSchema : SignUpFirstPageSchema),
  });

  const toggleFormType = () => {
    form.resetField("password");
    form.resetField("confirmPassword");
    form.clearErrors();
    setBEErrorMessage(undefined);
    setIsSignUp((prev) => !prev);
  };

  const postLogin = (authToken: string) => {
    // setToken(authToken);
    // navigation.navigate(); // Adjust based on your navigator
  };

  const handleLoginErr = async (error: any) => {
    const apiError = await error; // replace with ConvertToApiError if you have native API handler
    console.log(apiError);
    // setBEErrorMessage(TranslateApiErrorMessage(apiError));
  };

  const onSubmit = async (data: LoginOrSignUpPageFormType) => {
    // setLoading(true);
    try {
      if (isSignUp) {
        const _data = data as SignUpPageFormType;
        const res = await AuthController.apiV1PublicAuthSignupPost({
          githubComTDiblikProjectTemplateApiHandlersSignUpHandlerRequestBody: {
            email: _data.email,
            password: _data.password,
            useUsername: _data.useUsername,
            firstName: _data.firstName,
            lastName: _data.lastName,
            username: _data.username,
          },
        });
        postLogin(res.authToken);
      } else {
        const res = await AuthController.apiV1PublicAuthLoginPost({
          githubComTDiblikProjectTemplateApiHandlersLoginHandlerRequestBody: {
            email: data.email,
            password: data.password,
          },
        });
        postLogin(res.authToken);
      }
    } catch (err) {
      await handleLoginErr(err);
    } finally {
      // setLoading(false);
    }
  };

  return (
    <KeyboardAvoidingView behavior={Platform.OS === "ios" ? "padding" : undefined} className="flex-1 bg-base-200 justify-center px-4">
      <ScrollView contentContainerStyle={{flexGrow: 1, justifyContent: "center"}}>
        <View className="bg-base-100 p-8 rounded-2xl shadow-md">
          <Text className="text-3xl font-bold text-center mb-3">{isSignUp ? t("loginPage.signUpTitle") : t("loginPage.loginTitle")}</Text>

          <FormProvider {...form}>
            <View>
              <TextInput form={form} label={t("loginPage.email.label")} name="email" placeholder={t("loginPage.email.placeholder")} hasBigText />
              <PasswordFields t={t} form={form} isSignUp={isSignUp} />
              {beErrorMessage && <Text className="text-red-500 text-sm text-center mt-2">{beErrorMessage}</Text>}
              <TouchableOpacity className="bg-primary py-3 rounded-md mt-5 shadow-md" onPress={form.handleSubmit(onSubmit)}>
                <Text className="text-white text-center font-bold">{isSignUp ? t("loginPage.signUpButton") : t("loginPage.loginContinue")}</Text>
              </TouchableOpacity>
            </View>
          </FormProvider>

          <View className="flex-row items-center my-6">
            <View className="flex-1 h-px bg-gray-300" />
            <Text className="px-2 text-gray-500">{t("loginPage.continueWith")}</Text>
            <View className="flex-1 h-px bg-gray-300" />
          </View>

          <View className="grid grid-cols-2 gap-3 mb-6">
            <OAuthButton provider="Google" iconName={"google"} onPress={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectGoogleGet()} />
            <OAuthButton
              provider="Facebook"
              iconName={"facebook"}
              onPress={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectFacebookGet()}
            />
            <OAuthButton provider="Spotify" iconName={"spotify"} onPress={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectSpotifyGet()} />
            <OAuthButton provider="GitHub" iconName={"github"} onPress={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectGithubGet()} />
          </View>

          <Text className="text-base text-center mt-2">
            {isSignUp ? t("loginPage.signUpAlreadyHaveAccount") : t("loginPage.loginDontHaveAccount")}{" "}
            <Text onPress={toggleFormType} className="text-primary font-medium underline">
              {isSignUp ? t("loginPage.signUpLoginHere") : t("loginPage.loginSignUpHere")}
            </Text>
          </Text>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
};

const PasswordFields = ({t, form, isSignUp}: {t: any; form: any; isSignUp: boolean}) => (
  <View className="flex-row gap-4 mt-2">
    <TextInput
      form={form}
      label={t("loginPage.password.label")}
      name="password"
      placeholder={t("loginPage.password.placeholder")}
      hasBigText
      inputProps={{secureTextEntry: true}}
    />
  </View>
);

const OAuthButton = ({provider, iconName, onPress}: {provider: string; iconName: string; onPress: () => void}) => (
  <TouchableOpacity className="border border-gray-300 py-2 rounded-md flex-row items-center justify-center gap-2" onPress={onPress}>
    <FontAwesome6 name={iconName} size={24} color="black" />
    <Text>{provider}</Text>
  </TouchableOpacity>
);

export default LoginScreen;
