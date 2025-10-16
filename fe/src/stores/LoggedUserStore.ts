import type {UserModel} from "@shared/api-client";
import {useQuery, useQueryClient} from "@tanstack/react-query";
import {create} from "zustand";
import {UserController} from "../utils/api";
import {useEffect} from "react";

interface LoggedUserStoreState {
  user: UserModel | undefined;
  setUser: (user: UserModel | undefined) => void;
}

export const useLoggedUserStore = create<LoggedUserStoreState>((set) => ({
  user: undefined,
  setUser: (user) => set({user}),
}));

export const useFetchLoggedUser = () => {
  const {setUser} = useLoggedUserStore();

  const query = useQuery({
    queryKey: ["logged-user"],
    queryFn: () => UserController.apiV1PrivateUserMeGet(),
    retry: 10,
    retryDelay: 1_500,
  });

  useEffect(() => {
    if (query.data) setUser(query.data.userInfo);
    if (query.error) setUser(undefined);
  }, [query.data, query.error, setUser]);

  return query;
};

export const useLoggedUser = () => {
  const loggedUser = useLoggedUserStore((s) => s.user);
  const queryClient = useQueryClient();
  const refetchUser = () => queryClient.invalidateQueries({queryKey: ["logged-user"]});
  return {loggedUser, refetchUser};
};
