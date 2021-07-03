import { useQuery } from "react-query";
import { Channel } from "@liveassist/liber";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useChannels = (initialChannels?: Channel[]) => {
  const { isLoading, data, isError } = useQuery(
    QueryKeys.channels,
    () => api.channels.list(),
    { initialData: initialChannels }
  );
  return { isLoading, data, isError };
};
