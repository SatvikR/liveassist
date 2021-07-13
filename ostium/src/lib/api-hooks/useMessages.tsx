import { useEffect, useRef, useState } from "react";
import { Message, MessagingClient } from "@liveassist/liber";
import { useRefreshToken } from "./useRefreshToken";
import { MSG_URL } from "../api";
import { AccessToken } from "../AccessToken";

export const useMessages = (
  id: string,
  onMessage: (message: Message) => any
) => {
  const refreshToken = useRefreshToken();
  const [isConnecting, setConnecting] = useState<boolean>(true);
  const client = useRef<MessagingClient>(null);

  useEffect(() => {
    client.current = new MessagingClient(
      onMessage,
      () => {
        setConnecting(false);
      },
      MSG_URL
    );

    refreshToken().then(() =>
      client.current.connect(id, AccessToken.getInstance().value)
    );
  }, []);

  return { isConnecting, client };
};
