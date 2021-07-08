import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { useLoggedIn } from "../state/useLoggedIn";
import { QueryKeys } from "./keys";

export const useLogin = () => {
  const setLoginStatus = useLoggedIn((state) => state.setStatus);
  const queryClient = useQueryClient();
  const loginMutation = async (username: string, password: string) => {
    queryClient.cancelQueries(QueryKeys.me);
    const data = await api.users.login(username, password);

    if (!data.errors) {
      AccessToken.getInstance().value = data.accessToken;
      queryClient.invalidateQueries(QueryKeys.me);
      setLoginStatus(true);
    }

    return data.errors;
  };

  return loginMutation;
};
