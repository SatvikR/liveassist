import { Heading } from "@chakra-ui/layout";
import { Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { Channel } from "@liveassist/liber";
import React from "react";
import { useQuery } from "react-query";
import { api } from "../../lib/api";

export interface ChannelsProps {
  channels: Channel[];
}

export const Channels: React.FC<ChannelsProps> = ({ channels }) => {
  const { isLoading, data, isError } = useQuery("channels", api.channels.list, {
    initialData: channels,
  });

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
          <Box key={e.id} borderWidth="1px" p={4} my={6}>
            <Flex>
              <Text fontSize="large" color="gray.300" mr={3}>
                #
              </Text>
              <Text fontSize="large" isTruncated>
                {e.name}
              </Text>
            </Flex>
            <Flex mt={2}>
              {e.keywords.map((e, i) => (
                <Box
                  key={i}
                  mx={2}
                  bg="purple.300"
                  px={8}
                  py={1}
                  borderRadius={45}
                >
                  <Text>{e}</Text>
                </Box>
              ))}
            </Flex>
          </Box>
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
