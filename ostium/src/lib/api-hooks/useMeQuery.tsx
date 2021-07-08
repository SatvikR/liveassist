import { useQuery } from "react-query";
import { MeResponse } from "@liveassist/liber";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";
import { useLoggedIn } from "../state/useLoggedIn";

export const useMeQuery = (): {
  isLoading: boolean;
  data: MeResponse | null;
  isError: boolean;
} => {
  const accTok = AccessToken.getInstance();
  const setLoginStatus = useLoggedIn((state) => state.setStatus);

  const { isLoading, data, isError } = useQuery(
    QueryKeys.me,
    async () => {
      if (accTok.isExp()) {
        const ntok = await api.tokens.refresh();
        if (!ntok) {
          setLoginStatus(false);
          return null;
        }
        accTok.value = ntok;
        setLoginStatus(true);
      }
      return await api.users.me(accTok.value);
    },
    {
      refetchInterval: false,
      refetchOnMount: false,
    }
  );

  return { isLoading, data, isError };
};
