import {FaGithub, FaFacebook, FaGoogle, FaSpotify} from "react-icons/fa";
import {oAuthRedirectController} from "../../utils/api";
import {TextInput} from "../../components/TextInput";
import {FormProvider, useForm} from "react-hook-form";
import {LoginFirstPageSchema, zodResolver, type LoginFirstPageFormType} from "../../utils/validations";

export default function Login() {
  const form = useForm<LoginFirstPageFormType>({
    mode: "onTouched",
    resolver: zodResolver(LoginFirstPageSchema),
  });

  const onSubmit = (data: LoginFirstPageFormType) => {
    console.log(data);
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-base-200">
      <div className="card w-full max-w-md shadow-xl bg-base-100 p-8">
        <h2 className="text-2xl font-bold text-center mb-6">Login or Sign up</h2>

        <FormProvider {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <TextInput label={"Email"} name={"email"} placeholder="Enter your email" />
            <button className="btn btn-primary w-full mt-4" type="submit">
              Continue
            </button>
          </form>
        </FormProvider>

        <div className="divider">Or continue with</div>

        <div className="grid grid-cols-2 gap-3">
          <button
            className="btn btn-outline w-full flex items-center gap-2"
            onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectGoogleGet().then((s) => (window.location.href = s.redirectUrl!))}
          >
            <FaGoogle /> Google
          </button>
          <button
            className="btn btn-outline w-full flex items-center gap-2"
            onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectFacebookGet().then((s) => (window.location.href = s.redirectUrl!))}
          >
            <FaFacebook /> Facebook
          </button>
          <button
            className="btn btn-outline w-full flex items-center gap-2"
            onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectSpotifyGet().then((s) => (window.location.href = s.redirectUrl!))}
          >
            <FaSpotify /> Spotify
          </button>
          <button
            className="btn btn-outline w-full flex items-center gap-2"
            onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectGithubGet().then((s) => (window.location.href = s.redirectUrl!))}
          >
            <FaGithub /> GitHub
          </button>
        </div>
      </div>
    </div>
  );
}
