import React, {useEffect, useState} from "react";
import Layout from "../../components/Layout";
import {TextInput} from "../../components/TextInput";
import {FormProvider, useForm} from "react-hook-form";
import {zodResolver, SignUpFirstPageSchema, type SignUpPageFormType} from "../../utils/validations";
import {FaGithub, FaFacebook, FaGoogle, FaSpotify} from "react-icons/fa";
import {useLoadingStore} from "../../stores/LoadingStore";
import {t} from "i18next";
import {type GithubComTDiblikProjectTemplateApiModelsUsersModelDB} from "@shared/api-client";
import {oAuthRedirectController, UserController} from "../../utils/api";

export default function SettingsPage() {
  const [user, setUser] = useState<GithubComTDiblikProjectTemplateApiModelsUsersModelDB | null>(null);
  const {setLoading} = useLoadingStore();
  const form = useForm<SignUpPageFormType>({
    mode: "onChange",
    resolver: zodResolver(SignUpFirstPageSchema),
  });

  useEffect(() => {
    // todo: use some kind of user store that fetches the data when neccessary, using react-query
    UserController.apiV1PrivateUserMeGet().then((res) => {
      const info = res.userInfo;
      setUser(info ?? null);
      form.reset({
        email: info?.email ?? "",
        firstName: info?.firstName ?? "",
        lastName: info?.lastName ?? "",
        username: info?.handle ?? "",
        useUsername: !!info?.handle,
        password: "",
        confirmPassword: "",
      });
    });
  }, []);

  // todo: create and use a PATCH endpoint
  const onSubmit = async (data: SignUpPageFormType) => {
    setLoading(true);
    try {
      console.log("Update user:", data);
      alert("Profile updated!");
    } finally {
      setLoading(false);
    }
  };

  // todo: extract into a separate type and out of the render loop
  // todo: add an "x" that can redact the oauth connection
  // todo: fix the problem that when connecting different accounts with different emails, it just fucks everything
  const OAuthButton = ({
    provider,
    icon,
    connected,
    textConnect,
    textConnected,
    onConnect,
  }: {
    provider: string;
    icon: React.ReactNode;
    connected?: boolean;
    textConnect: string;
    textConnected: string;
    onConnect: () => Promise<any>;
  }) => (
    <button
      className={`flex items-center gap-2 px-4 py-2 rounded-md w-full justify-start ${
        connected ? "bg-green-100 text-green-700 cursor-not-allowed" : "border border-gray-300 hover:bg-gray-50 cursor-pointer"
      }`}
      onClick={() => !connected && onConnect().then((s) => (window.location.href = s.redirectUrl!))}
    >
      {icon} {provider} {connected ? textConnected : textConnect}
    </button>
  );

  return (
    <Layout>
      <div className="max-w-4xl mx-auto py-10 space-y-8">
        <h1 className="text-3xl font-bold">{t("settingsPage.pageTitle")}</h1>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {/* Left Column: Avatar + OAuth */}
          <div className="space-y-6">
            <div className="flex flex-col items-center p-6 rounded-2xl shadow-md bg-base-100">
              <img
                src={user?.avatarUrl ?? "/default-avatar.png"}
                alt={user?.userFullName ?? "Avatar"}
                className="w-32 h-32 rounded-full object-cover mb-4"
              />
              <button className="btn btn-outline w-full">{t("settingsPage.changeAvatar")}</button>
            </div>

            <div className="p-6 rounded-2xl shadow-md bg-base-100 space-y-3">
              <h2 className="font-semibold">{t("settingsPage.connectedAccounts")}</h2>
              <div className="space-y-2">
                <OAuthButton
                  provider="Google"
                  icon={<FaGoogle />}
                  connected={!!user?.googleId}
                  textConnect={t("settingsPage.oauth.connect")}
                  textConnected={t("settingsPage.oauth.connected")}
                  onConnect={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectGoogleGet({redirectBackToAfterOauth: "settings"})}
                />
                <OAuthButton
                  provider="Facebook"
                  icon={<FaFacebook />}
                  connected={!!user?.facebookId}
                  textConnect={t("settingsPage.oauth.connect")}
                  textConnected={t("settingsPage.oauth.connected")}
                  onConnect={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectFacebookGet({redirectBackToAfterOauth: "settings"})}
                />
                <OAuthButton
                  provider="Spotify"
                  icon={<FaSpotify />}
                  connected={!!user?.spotifyId}
                  textConnect={t("settingsPage.oauth.connect")}
                  textConnected={t("settingsPage.oauth.connected")}
                  onConnect={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectSpotifyGet({redirectBackToAfterOauth: "settings"})}
                />
                <OAuthButton
                  provider="Github"
                  icon={<FaGithub />}
                  connected={!!user?.githubHandle}
                  textConnect={t("settingsPage.oauth.connect")}
                  textConnected={t("settingsPage.oauth.connected")}
                  onConnect={() => oAuthRedirectController.apiV1PublicAuthOauthRedirectGithubGet({redirectBackToAfterOauth: "settings"})}
                />
              </div>
            </div>
          </div>

          {/* Right Column: Editable Fields */}
          <div className="md:col-span-2 space-y-6">
            <FormProvider {...form}>
              <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
                <div className="p-6 rounded-2xl shadow-md bg-base-100 space-y-4">
                  <h2 className="font-semibold text-lg">{t("settingsPage.profileInfo")}</h2>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <TextInput
                      label={t("settingsPage.firstName.label")}
                      name="firstName"
                      placeholder={t("settingsPage.firstName.placeholder")}
                      hasBigText
                    />
                    <TextInput
                      label={t("settingsPage.lastName.label")}
                      name="lastName"
                      placeholder={t("settingsPage.lastName.placeholder")}
                      hasBigText
                    />
                  </div>
                  <TextInput label={t("settingsPage.email.label")} name="email" placeholder={t("settingsPage.email.placeholder")} hasBigText />
                  <TextInput
                    label={t("settingsPage.username.label")}
                    name="username"
                    placeholder={t("settingsPage.username.placeholder")}
                    hasBigText
                  />
                  <button type="submit" className="btn btn-primary w-full mt-4">
                    {t("settingsPage.saveChanges")}
                  </button>
                </div>

                <div className="p-6 rounded-2xl shadow-md bg-base-100 space-y-4">
                  <h2 className="font-semibold text-lg">{t("settingsPage.changePassword")}</h2>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <TextInput
                      label={t("settingsPage.password.label")}
                      name="password"
                      placeholder={t("settingsPage.password.placeholder")}
                      inputProps={{type: "password"}}
                      hasBigText
                    />
                    <TextInput
                      label={t("settingsPage.confirmPassword.label")}
                      name="confirmPassword"
                      placeholder={t("settingsPage.confirmPassword.placeholder")}
                      inputProps={{type: "password"}}
                      hasBigText
                    />
                  </div>
                  <button type="submit" className="btn btn-primary w-full mt-4">
                    {t("settingsPage.savePassword")}
                  </button>
                </div>
              </form>
            </FormProvider>
          </div>
        </div>
      </div>
    </Layout>
  );
}
