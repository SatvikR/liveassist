import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { useLoggedIn } from "../state/useLoggedIn";
import { QueryKeys } from "./keys";

export const useLogout = () => {
  const setLoginStatus = useLoggedIn((state) => state.setStatus);
  const queryClient = useQueryClient();
  const logoutMutation = async () => {
    queryClient.cancelQueries(QueryKeys.me);
    await api.users.logout();
    AccessToken.getInstance().reset();
    queryClient.invalidateQueries(QueryKeys.me);
    setLoginStatus(false);
  };

  return logoutMutation;
};
