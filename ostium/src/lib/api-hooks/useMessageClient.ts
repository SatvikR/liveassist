import { Message, MessagingClient } from "@liveassist/liber";
import { useEffect, useRef, useState } from "react";
import { AccessToken } from "../AccessToken";
import { MSG_URL } from "../api";
import { useRefreshToken } from "./useRefreshToken";

export const useMessageClient = (
  id: string,
  onMessage: (message: Message) => any,
  onMessages: (messages: Message[]) => any
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
      onMessages,
      MSG_URL
    );

    refreshToken().then(() =>
      client.current.connect(id, AccessToken.getInstance().value)
    );
  }, []);

  return { isConnecting, client };
};
