import { Heading } from "@chakra-ui/layout";
import { Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { Channel as IChannel } from "@liveassist/liber";
import React from "react";
import { useChannels } from "../../lib/api-hooks/useChannels";
import { Channel } from "./Channel";

export interface ChannelsProps {
  channels: IChannel[];
}

export const Channels: React.FC<ChannelsProps> = ({ channels }) => {
  const { isLoading, data, isError } = useChannels(channels);

  let body: JSX.Element;
  if (isLoading) {
    body = (
      <Spinner
        thickness="4px"
        speed="0.5s"
        emptyColor="gray.200"
        color="blue.300"
        size="xl"
      />
    );
  } else if (isError) {
    body = <Text>Error... </Text>;
  } else
    body = (
      <Box my={12}>
        {data.map((e) => (
          <Channel channel={e} key={e.id} />
        ))}
      </Box>
    );

  return (
    <>
      <Heading>Channels</Heading>
      {body}
    </>
  );
};
