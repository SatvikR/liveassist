import { AccessToken } from "../AccessToken";
import { api } from "../api";
import { useLoggedIn } from "../state/useLoggedIn";

export const useRefreshToken = () => {
  const accTok = AccessToken.getInstance();
  const setLoginStatus = useLoggedIn((state) => state.setStatus);

  const refreshToken = async () => {
    if (accTok.isExp()) {
      const ntok = await api.tokens.refresh();
      if (!ntok) {
        setLoginStatus(false);
        return;
      }
      accTok.value = ntok;
      setLoginStatus(true);
    }
  };

  return refreshToken;
};
