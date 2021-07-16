import { Box, Heading, Spinner } from "@chakra-ui/react";
import React from "react";
import { useMeQuery } from "../../lib/api-hooks/useMeQuery";

export interface UserDataProps {}

export const UserData: React.FC<UserDataProps> = ({}) => {
  const { isLoading, data, isError } = useMeQuery();

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
    body = <Heading>Hello there, {data.username}</Heading>;
  }

  return <Box>{body}</Box>;
};
