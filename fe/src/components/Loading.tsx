import {motion} from "motion/react";
import {useTranslation} from "react-i18next";

export type LoaderTextPossibilities = undefined | "loadingStates.default" | "loadingStates.loggingIn";
export const Loader = ({textCode}: {textCode?: LoaderTextPossibilities}) => {
  const {t} = useTranslation();
  return (
    <motion.div
      className="fixed inset-0 flex items-center justify-center bg-black/30 z-50"
      initial={{opacity: 0}}
      animate={{opacity: 1}}
      exit={{opacity: 0}}
    >
      <div className="flex flex-col items-center gap-2 bg-base-100 p-6 rounded-xl shadow-lg">
        <div className="animate-spin rounded-full h-12 w-12 border-t-4 border-primary border-solid"></div>
        <p className="text-lg font-medium text-center mt-1.5">{t(textCode ?? "loadingStates.default")}</p>
      </div>
    </motion.div>
  );
};
