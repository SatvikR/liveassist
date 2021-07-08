import { useQuery } from "react-query";
import { MeResponse } from "@liveassist/liber";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useMeQuery = (): {
  isLoading: boolean;
  data: MeResponse | null;
  isError: boolean;
} => {
  const accTok = AccessToken.getInstance();

  const { isLoading, data, isError } = useQuery(
    QueryKeys.me,
    async () => {
      if (accTok.isExp()) {
        const ntok = await api.tokens.refresh();
        if (!ntok) {
          return null;
        }
        accTok.value = ntok;
      }
      return await api.users.me(accTok.value);
    },
    {
      refetchInterval: false,
    }
  );

  return { isLoading, data, isError };
};
