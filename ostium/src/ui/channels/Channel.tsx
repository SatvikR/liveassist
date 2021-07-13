import { Box, Flex, Text } from "@chakra-ui/react";
import { Channel as IChannel } from "@liveassist/liber";
import React from "react";

export interface ChannelProps {
  channel: IChannel;
}

export const Channel: React.FC<ChannelProps> = ({ channel }) => {
  return (
    <Box borderWidth="1px" p={4} my={6}>
      <Flex>
        <Text fontSize="large" color="gray.300" mr={3}>
          #
        </Text>
        <Text fontSize="large" isTruncated>
          {channel.name}
        </Text>
      </Flex>
      <Flex mt={2}>
        {channel.keywords.map((e, i) => (
          <Box key={i} mx={2} bg="purple.300" px={8} py={1} borderRadius={45}>
            <Text>{e}</Text>
          </Box>
        ))}
      </Flex>
    </Box>
  );
};
