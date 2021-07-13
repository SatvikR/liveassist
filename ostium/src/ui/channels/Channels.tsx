import { Heading } from "@chakra-ui/layout";
import { Box, chakra, Flex, Spinner, Text } from "@chakra-ui/react";
import { Channel as IChannel } from "@liveassist/liber";
import NextLink from "next/link";
import React from "react";
import { StyledButton } from "../../components/StyledButton";
import { useChannels } from "../../lib/api-hooks/useChannels";
import { Channel } from "./Channel";
import { useRouter } from "next/router";

export interface ChannelsProps {
  channels: IChannel[];
}

export const Channels: React.FC<ChannelsProps> = ({ channels }) => {
  const { isLoading, data, isError } = useChannels(channels);
  const router = useRouter();

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
          <NextLink href={`/channel/${e.id}`} key={e.id}>
            <chakra.a
              textDecor="none"
              color="black"
              _hover={{ cursor: "pointer" }}
            >
              <Channel channel={e} />
            </chakra.a>
          </NextLink>
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
