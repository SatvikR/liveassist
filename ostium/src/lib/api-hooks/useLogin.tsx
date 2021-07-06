import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useLogin = () => {
  const queryClient = useQueryClient();
  const loginMutation = async (username: string, password: string) => {
    queryClient.cancelQueries(QueryKeys.me);
    const data = await api.users.login(username, password);

    if (!data.errors) {
      AccessToken.value = data.accessToken;
      queryClient.invalidateQueries(QueryKeys.me);
    }

    return data.errors;
  };

  return loginMutation;
};
