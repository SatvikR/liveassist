import {
  Box,
  Flex,
  Heading,
  Input,
  Spinner,
  Stack,
  Text,
} from "@chakra-ui/react";
import { Message } from "@liveassist/liber";
import React, { useRef, useState } from "react";
import { StyledButton } from "../../components/StyledButton";
import { useChannel } from "../../lib/api-hooks/useChannel";
import { useMessageClient } from "../../lib/api-hooks/useMessageClient";

export interface ChatProps {
  id: string;
}

export const Chat: React.FC<ChatProps> = ({ id }) => {
  if (!id) {
    return null;
  }

  const [messages, setMessages] = useState<Message[]>([]);
  const [msg, setMsg] = useState<string>("");
  const bottomMsg = useRef<HTMLDivElement>(null);

  const { isLoading, data, isError } = useChannel(id);
  const { isConnecting, client } = useMessageClient(id, (m) => {
    setMessages((oldMessages) => [...oldMessages, m]);
    bottomMsg.current.scrollIntoView({ behavior: "smooth" });
  });

  if (isLoading || isError) {
    return (
      <Box>
        <Spinner
          thickness="4px"
          speed="0.65s"
          emptyColor="gray.200"
          color="blue.500"
          size="xl"
        />
      </Box>
    );
  }

  return (
    <>
      {isConnecting ? (
        <Box>
          <Heading>Connecting...</Heading>
        </Box>
      ) : (
        <Box>
          <Box mb={4}>
            <Heading>{data.name}</Heading>
          </Box>
          <Stack overflowY="auto" height="75vh" spacing={4}>
            {messages.map((e, i) => {
              const isFirst = i == 0 ? { mt: "auto" } : {};

              return (
                <Box borderTopWidth="1px" key={i} {...isFirst}>
                  <Flex>
                    <Box>
                      <Text fontWeight="bold">{e.owner.username}</Text>
                    </Box>
                    <Box ml={8}>
                      <Text>{new Date(e.createdAt).toLocaleString()}</Text>
                    </Box>
                  </Flex>
                  <Text>{e.text}</Text>
                </Box>
              );
            })}
            <Box ref={bottomMsg}></Box>
          </Stack>
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
        </Box>
      )}
    </>
  );
};
