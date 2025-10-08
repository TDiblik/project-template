import {AnimatePresence} from "motion/react";
import {useLoadingStore} from "../stores/LoadingStore";
import {Loader} from "./Loading";

export const LoaderProvider: React.FC<React.PropsWithChildren> = ({children}) => {
  const {loading, loadingTextCode} = useLoadingStore();

  return (
    <>
      <AnimatePresence>{loading && <Loader textCode={loadingTextCode} />}</AnimatePresence>
      {children}
    </>
  );
};
