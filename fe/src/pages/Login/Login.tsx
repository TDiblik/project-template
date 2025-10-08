import {useState} from "react";
import {AnimatePresence, motion, type HTMLMotionProps} from "motion/react";
import {FaGithub, FaFacebook, FaGoogle, FaSpotify} from "react-icons/fa";
import {AuthController, ConvertToApiError, oAuthRedirectController} from "../../utils/api";
import {TextInput} from "../../components/TextInput";
import {FormProvider, useForm} from "react-hook-form";
import {
  LoginFirstPageSchema,
  SignUpFirstPageSchema,
  zodResolver,
  type LoginOrSignUpPageFormType,
  type SignUpPageFormType,
} from "../../utils/validations";
import {HiddenBooleanInput} from "../../components/HiddenBooleanInput";
import {useFormLog} from "../../utils/useFormLog";

// todo: redirect after successfull signup/login
// todo: login
// todo: handle errors
// todo: show loader while fetching

export default function Login() {
  const [isSignUp, setIsSignUp] = useState(false);

  const form = useForm<LoginOrSignUpPageFormType>({
    mode: "onChange",
    resolver: zodResolver(!isSignUp ? LoginFirstPageSchema : SignUpFirstPageSchema),
  });
  useFormLog(form);

  const toggleFormType = () => {
    form.clearErrors();
    setIsSignUp((prev) => !prev);
  };

  const onSubmit = async (data: LoginOrSignUpPageFormType) => {
    if (isSignUp) {
      const _data = data as SignUpPageFormType;
      AuthController.apiV1AuthSignupPost({
        githubComTDiblikProjectTemplateApiHandlersSignUpRequestBody: {
          email: _data.email,
          password: _data.password,
          useUsername: _data.useUsername,
          firstName: _data.firstName,
          lastName: _data.lastName,
          username: _data.username,
        },
      })
        .then((s) => {
          console.log(s);
        })
        .catch(async (err) => {
          const erorr = await ConvertToApiError(err);
          console.log(erorr);
        });
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-base-200 px-4">
      <motion.div
        layout
        transition={{type: "spring", stiffness: 120, damping: 20}}
        className="card w-full max-w-md shadow-2xl bg-base-100 p-8 rounded-2xl"
      >
        <motion.h2
          key={isSignUp ? "sign-up" : "login"}
          initial={{opacity: 0, y: -15}}
          animate={{opacity: 1, y: 0}}
          exit={{opacity: 0, y: 15}}
          transition={{duration: 0.3}}
          className="text-3xl font-bold text-center mb-3"
        >
          {isSignUp ? "Create your account" : "Welcome back"}
        </motion.h2>

        <FormProvider {...form}>
          <AnimatePresence mode="wait">
            <motion.form
              key={isSignUp ? "signup-form" : "login-form"}
              initial={{opacity: 0, y: 15}}
              animate={{opacity: 1, y: 0}}
              exit={{opacity: 0, y: -15}}
              transition={{duration: 0.35}}
              onSubmit={form.handleSubmit(onSubmit)}
            >
              {isSignUp && <NameFields />}
              <TextInput label="Email" name="email" placeholder="Enter your email" inputProps={{type: "email"}} hasBigText />
              <PasswordFields isSignUp={isSignUp} />
              <motion.button
                whileHover={{scale: 1.03}}
                whileTap={{scale: 0.97}}
                className="btn btn-primary w-full mt-5 border-0 text-white shadow-md hover:shadow-lg transition-all"
                type="submit"
              >
                {isSignUp ? "Sign up" : "Continue"}
              </motion.button>
            </motion.form>
          </AnimatePresence>
        </FormProvider>

        <div className="divider my-6">Or continue with</div>

        <div className="grid grid-cols-2 gap-3 mb-6">
          <OAuthButton provider="Google" icon={<FaGoogle />} onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectGoogleGet()} />
          <OAuthButton provider="Facebook" icon={<FaFacebook />} onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectFacebookGet()} />
          <OAuthButton provider="Spotify" icon={<FaSpotify />} onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectSpotifyGet()} />
          <OAuthButton provider="GitHub" icon={<FaGithub />} onClick={() => oAuthRedirectController.apiV1AuthOauthRedirectGithubGet()} />
        </div>

        <motion.p className="text-base text-center" initial={{opacity: 0}} animate={{opacity: 1}} transition={{delay: 0.2}}>
          {isSignUp ? "Already have an account?" : "Don't have an account?"}{" "}
          <button
            type="button"
            onClick={toggleFormType}
            className="text-primary font-medium underline hover:text-secondary transition-colors cursor-pointer"
          >
            {isSignUp ? "Login here" : "Sign up here"}
          </button>
        </motion.p>
      </motion.div>
    </div>
  );
}

export const NameFields = () => {
  const [useUsername, setUseUsername] = useState(false);
  const animation = {
    initial: {opacity: 0, scale: 0.95},
    animate: {opacity: 1, scale: 1},
    exit: {opacity: 0, scale: 0.95},
    transition: {duration: 0.25},
  } as HTMLMotionProps<"div">;

  return (
    <div className="flex flex-col gap-1">
      <div className="flex justify-center">
        <button
          type="button"
          onClick={() => setUseUsername(false)}
          className={`btn btn-sm rounded-l-full ${!useUsername ? "btn-primary text-white" : "btn-outline"}`}
        >
          Use Name
        </button>
        <button
          type="button"
          onClick={() => setUseUsername(true)}
          className={`btn btn-sm rounded-r-full ${useUsername ? "btn-primary text-white" : "btn-outline"}`}
        >
          Use Username
        </button>
      </div>
      {!useUsername ? (
        <motion.div key="name-fields" {...animation} className="flex gap-4">
          <TextInput label="First Name" name="firstName" placeholder="Enter your first name" inputProps={{type: "text"}} hasBigText />
          <TextInput label="Last Name" name="lastName" placeholder="Enter your last name" inputProps={{type: "text"}} hasBigText />
        </motion.div>
      ) : (
        <motion.div key="username-field" {...animation}>
          <TextInput label="Username" name="username" placeholder="Enter your username" inputProps={{type: "text"}} hasBigText />
        </motion.div>
      )}
      <HiddenBooleanInput name="useUsername" value={useUsername} />
    </div>
  );
};
const PasswordFields = ({isSignUp}: {isSignUp: boolean}) => (
  <>
    <div className="flex gap-4">
      <TextInput label="Password" name="password" placeholder="Enter your password" inputProps={{type: "password"}} hasBigText />
      {isSignUp && (
        <TextInput label="Confirm Password" name="confirmPassword" placeholder="Re-enter your password" inputProps={{type: "password"}} hasBigText />
      )}
    </div>
  </>
);

const OAuthButton = ({provider, icon, onClick}: {provider: string; icon: React.ReactNode; onClick: () => Promise<any>}) => (
  <motion.button
    whileHover={{scale: 1.05}}
    whileTap={{scale: 0.97}}
    className="btn btn-outline w-full flex items-center gap-2 transition-all hover:border-primary"
    onClick={() => onClick().then((s) => (window.location.href = s.redirectUrl!))}
  >
    {icon} {provider}
  </motion.button>
);
