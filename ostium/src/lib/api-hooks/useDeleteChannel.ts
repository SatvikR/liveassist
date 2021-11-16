import { QueryKeys } from "./keys";
import { useQueryClient } from "react-query";
import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { useRefreshToken } from "./useRefreshToken";

export const useDeleteChannel = () => {
  const queryClient = useQueryClient();
  const refreshToken = useRefreshToken();

  const deleteChannel = async (id: string) => {
    await refreshToken();
    await api.channels.delete(id, AccessToken.getInstance().value);
    queryClient.invalidateQueries(QueryKeys.channels);
    queryClient.invalidateQueries(
      QueryKeys.userChannels(AccessToken.getInstance().userId)
    );
  };

  return deleteChannel;
};
