import { Channel } from "@liveassist/liber";
import { useQuery } from "react-query";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useChannel = (id: string, initial: Channel) => {
  const { isLoading, data, isError } = useQuery(
    QueryKeys.channel(id),
    () => api.channels.get(id),
    { initialData: initial }
  );
  return { isLoading, data, isError };
};
