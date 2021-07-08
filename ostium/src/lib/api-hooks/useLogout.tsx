import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useLogout = () => {
  const queryClient = useQueryClient();
  const logoutMutation = async () => {
    queryClient.cancelQueries(QueryKeys.me);
    await api.users.logout();
    AccessToken.getInstance().reset();
    queryClient.invalidateQueries(QueryKeys.me);
  };

  return logoutMutation;
};
