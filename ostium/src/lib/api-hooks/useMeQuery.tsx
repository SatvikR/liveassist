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
  if (!AccessToken.value) {
    return {
      isLoading: false,
      data: null,
      isError: false,
    };
  }
  const { isLoading, data, isError } = useQuery(
    QueryKeys.me,
    () => api.users.me(AccessToken.value),
    {
      refetchInterval: false,
    }
  );

  return { isLoading, data, isError };
};
