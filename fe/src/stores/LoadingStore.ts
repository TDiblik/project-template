import {create} from "zustand";

interface LoadingStoreState {
  loading: boolean;
  loadingText?: string;
  setLoading: (newState: boolean, loadingText?: string) => void;
}

export const useLoadingStore = create<LoadingStoreState>()((set) => ({
  loading: false,
  setLoading: (newState, loadingText) => set(() => ({loading: newState, loadingText})),
}));
