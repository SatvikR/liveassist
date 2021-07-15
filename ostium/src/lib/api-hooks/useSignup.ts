import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { useLoggedIn } from "../state/useLoggedIn";
import { QueryKeys } from "./keys";

export const useSignup = () => {
  const setLoginStatus = useLoggedIn((state) => state.setStatus);
  const queryClient = useQueryClient();
  const signupMutation = async (
    username: string,
    email: string,
    password: string
  ) => {
    queryClient.cancelQueries(QueryKeys.me);
    const data = await api.users.signup(username, email, password);

    if (!data.errors) {
      AccessToken.getInstance().value = data.accessToken;
      queryClient.invalidateQueries(QueryKeys.me);
      setLoginStatus(true);
    }

    return data.errors;
  };

  return signupMutation;
};
