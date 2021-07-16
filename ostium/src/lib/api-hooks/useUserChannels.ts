import { useQuery } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";

export const useUserChannels = () => {
  const uid = AccessToken.getInstance().userId;
  const { isLoading, data, isError } = useQuery(
    QueryKeys.userChannels(uid),
    () => api.channels.getByUser(uid)
  );

  return { isLoading, data, isError };
};
