import {
  Box,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Stack,
  Text,
} from "@chakra-ui/react";
import { Message } from "@liveassist/liber";
import React, { useEffect, useState } from "react";
import { StyledButton } from "../../components/StyledButton";
import { useMessages } from "../../lib/api-hooks/useMessages";

export interface ChatProps {
  id: string;
}

export const Chat: React.FC<ChatProps> = ({ id }) => {
  if (!id) {
    return null;
  }

  const [messages, setMessages] = useState<Message[]>([]);
  const [msg, setMsg] = useState<string>("");

  const { isConnecting, client } = useMessages(id, (m) => {
    setMessages((oldMessages) => [...oldMessages, m]);
  });

  return (
    <>
      {isConnecting ? (
        <Box>
          <Heading>Connecting...</Heading>
        </Box>
      ) : (
        <Box>
          <Heading>connected to {id}</Heading>
          <Flex my={2}>
            <Input
              placeholder="Send a message"
              value={msg}
              onChange={(e) => setMsg(e.target.value)}
            />
            <StyledButton
              ml={4}
              onClick={(_) => {
                setMsg("");
                client.current.send(msg);
              }}
            >
              Send
            </StyledButton>
          </Flex>
          <Stack>
            {messages.map((e, i) => (
              <Box key={i} p={4} my={4} borderWidth="1px">
                <Flex>
                  <Text>{e.owner.username}</Text>
                  <Box ml="auto">
                    <Text>{new Date(e.createdAt).toLocaleString()}</Text>
                  </Box>
                </Flex>
                <Text>{e.text}</Text>
              </Box>
            ))}
          </Stack>
        </Box>
      )}
    </>
  );
};
