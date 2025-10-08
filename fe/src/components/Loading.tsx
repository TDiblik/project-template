import {motion} from "motion/react";

export const Loader = ({text = "Loading..."}: {text?: string}) => (
  <motion.div
    className="fixed inset-0 flex items-center justify-center bg-black/30 z-50"
    initial={{opacity: 0}}
    animate={{opacity: 1}}
    exit={{opacity: 0}}
  >
    <div className="flex flex-col items-center gap-2 bg-base-100 p-6 rounded-xl shadow-lg">
      <div className="animate-spin rounded-full h-12 w-12 border-t-4 border-primary border-solid"></div>
      <p className="text-lg font-medium text-center">{text}</p>
    </div>
  </motion.div>
);
