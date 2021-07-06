import { api } from "../api";
import { useTokenStore } from "../state/useToken";
import { QueryKeys } from "./keys";
import { useQuery } from "react-query";

export const useMeQuery = () => {
  const token = useTokenStore((state) => state.token);
  const { isLoading, data, isError } = useQuery(QueryKeys.me, () =>
    api.users.me(token)
  );

  return { isLoading, data, isError };
};
