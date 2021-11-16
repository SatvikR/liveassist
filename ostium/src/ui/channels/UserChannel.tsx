import { Box, Flex, IconButton, Text } from "@chakra-ui/react";
import { Channel as IChannel } from "@liveassist/liber";
import React from "react";
import { DeleteIcon, EditIcon, ExternalLinkIcon } from "@chakra-ui/icons";
import { useRouter } from "next/router";
import { useDeleteChannel } from "../../lib/api-hooks/useDeleteChannel";

export interface UserChannelProps {
  channel: IChannel;
}

export const UserChannel: React.FC<UserChannelProps> = ({ channel }) => {
  const router = useRouter();
  const deleteChannel = useDeleteChannel();

  return (
    <Box borderWidth="1px" p={4} my={6}>
      <Flex>
        <Box>
          <IconButton
            aria-label="Delete channel"
            icon={<DeleteIcon />}
            colorScheme="red"
            borderRadius={5}
            size="xs"
            mx={1}
            onClick={() => deleteChannel(channel.id)}
          />
          <IconButton
            aria-label="Edit channel"
            icon={<EditIcon />}
            colorScheme="yellow"
            borderRadius={5}
            size="xs"
            mx={1}
          />
          <IconButton
            aria-label="View Channel"
            icon={<ExternalLinkIcon />}
            colorScheme="teal"
            borderRadius={5}
            size="xs"
            mx={1}
            mr={4}
            onClick={() => router.push(`/channel/${channel.id}`)}
          />
        </Box>
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
