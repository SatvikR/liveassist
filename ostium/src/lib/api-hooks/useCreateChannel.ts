import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { QueryKeys } from "./keys";
import { useRefreshToken } from "./useRefreshToken";

export const useCreateChannel = () => {
  const queryClient = useQueryClient();
  const refreshToken = useRefreshToken();

  const createChannel = async (title: string, keywords: string[]) => {
    await refreshToken();
    await api.channels.create(title, keywords, AccessToken.getInstance().value);
    queryClient.invalidateQueries(QueryKeys.channels);
  };

  return createChannel;
};
