import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useSignup = () => {
  const queryClient = useQueryClient();
  const signupMutation = async (
    username: string,
    email: string,
    password: string
  ) => {
    queryClient.cancelQueries(QueryKeys.me);
    const data = await api.users.signup(username, email, password);

    if (!data.errors) {
      AccessToken.value = data.accessToken;
      queryClient.invalidateQueries(QueryKeys.me);
    }

    return data.errors;
  };

  return signupMutation;
};
