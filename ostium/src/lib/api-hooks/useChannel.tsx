import { useQuery } from "react-query";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useChannel = (id: string) => {
  const { isLoading, data, isError } = useQuery(QueryKeys.channel(id), () =>
    api.channels.get(id)
  );
  return { isLoading, data, isError };
};
