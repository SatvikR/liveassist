import { Box, chakra, Spinner } from "@chakra-ui/react";
import NextLink from "next/link";
import React from "react";
import { useUserChannels } from "../../lib/api-hooks/useUserChannels";
import { UserChannel } from "../channels/UserChannel";
import { UserData } from "./UserData";

export interface ProfileProps {}

export const Profile: React.FC<ProfileProps> = ({}) => {
  const { isLoading, data, isError } = useUserChannels();

  let body: JSX.Element;

  if (isLoading || isError) {
    body = (
      <Spinner
        thickness="4px"
        speed="0.65s"
        emptyColor="gray.200"
        color="blue.500"
        size="xl"
      />
    );
  } else {
    body = (
      <>
        {data.map((e) => (
          <UserChannel channel={e} />
        ))}
      </>
    );
  }

  return (
    <Box>
      <UserData />
      <Box my={12}>{body}</Box>
    </Box>
  );
};
