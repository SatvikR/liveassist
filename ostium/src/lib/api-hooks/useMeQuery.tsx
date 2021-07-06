import { useQuery } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useMeQuery = () => {
  const { isLoading, data, isError } = useQuery(QueryKeys.me, () =>
    api.users.me(AccessToken.value)
  );

  return { isLoading, data, isError };
};
