import {create} from "zustand";
import type {LoaderTextPossibilities} from "../components/Loading";

interface LoadingStoreState {
  loading: boolean;
  loadingTextCode?: LoaderTextPossibilities;
  setLoading: (loading: boolean, loadingTextCode?: LoaderTextPossibilities) => void;
}

export const useLoadingStore = create<LoadingStoreState>()((set) => ({
  loading: false,
  setLoading: (loading, loadingTextCode) => set(() => ({loading, loadingTextCode})),
}));
