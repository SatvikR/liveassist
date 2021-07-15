import { MeResponse } from "@liveassist/liber";
import { useQuery } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";
import { useRefreshToken } from "./useRefreshToken";

export const useMeQuery = (): {
  isLoading: boolean;
  data: MeResponse | null;
  isError: boolean;
} => {
  const accTok = AccessToken.getInstance();
  const refreshToken = useRefreshToken();

  const { isLoading, data, isError } = useQuery(
    QueryKeys.me,
    async () => {
      await refreshToken();
      return await api.users.me(accTok.value);
    },
    {
      refetchInterval: false,
      refetchOnMount: false,
    }
  );

  return { isLoading, data, isError };
};
