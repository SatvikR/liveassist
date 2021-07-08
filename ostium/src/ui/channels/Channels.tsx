import { Heading } from "@chakra-ui/layout";
import { Box, Button, Flex, Spinner, Text } from "@chakra-ui/react";
import { Channel as IChannel } from "@liveassist/liber";
import React from "react";
import { useChannels } from "../../lib/api-hooks/useChannels";
import { Channel } from "./Channel";
import NextLink from "next/link";
import { StyledButton } from "../../components/StyledButton";
import { useRouter } from "next/dist/client/router";
import { useLoggedIn } from "../../lib/state/useLoggedIn";

export interface ChannelsProps {
  channels: IChannel[];
}

export const Channels: React.FC<ChannelsProps> = ({ channels }) => {
  const { isLoading, data, isError } = useChannels(channels);
  const router = useRouter();
  const loggedIn = useLoggedIn((state) => state.loggedIn);
  const loading = useLoggedIn((state) => state.loading);
  console.log(`loggedin: ${loggedIn}, loading: ${loading}`);

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
      <Flex>
        <Box>
          <Heading>Channels</Heading>
        </Box>
        <Box ml="auto">
          <StyledButton onClick={() => router.push("/create-channel")}>
            Create Channel
          </StyledButton>
        </Box>
      </Flex>
      {body}
    </>
  );
};
