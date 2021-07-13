import { Box, Flex, Text } from "@chakra-ui/react";
import { Message as IMessage } from "@liveassist/liber";
import React from "react";

export interface MessageProps {
  isFirst: boolean;
  message: IMessage;
}

export const Message: React.FC<MessageProps> = ({ isFirst, message }) => {
  const margin = isFirst ? { mt: "auto" } : {};

  return (
    <Box borderTopWidth="1px" {...margin}>
      <Flex>
        <Box>
          <Text fontWeight="bold">{message.owner.username}</Text>
        </Box>
        <Box ml={8}>
          <Text>{new Date(message.createdAt).toLocaleString()}</Text>
        </Box>
      </Flex>
      <Text>{message.text}</Text>
    </Box>
  );
};
