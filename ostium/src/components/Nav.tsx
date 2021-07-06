import { ArrowForwardIcon } from "@chakra-ui/icons";
import { Flex } from "@chakra-ui/layout";
import { Box, Link, Text } from "@chakra-ui/react";
import NextLink from "next/link";
import React from "react";
import { useMeQuery } from "../lib/api-hooks/useMeQuery";
import { gradient } from "./constants";

export interface NavProps {}

export const Nav: React.FC<NavProps> = ({}) => {
  const { isLoading, data, isError } = useMeQuery();

  return (
    <Flex top={0} p={4}>
      <Box ml={4}>
        <NextLink href="/">
          <Link
            fontSize="3xl"
            fontWeight="bold"
            bgGradient={gradient}
            bgClip="text"
          >
            LiveAssist
          </Link>
        </NextLink>
      </Box>
      <Box ml="auto">
        {isLoading || isError || !data ? (
          <NextLink href="/login">
            <Link mr={8}>
              Login <ArrowForwardIcon mb={0.5} />
            </Link>
          </NextLink>
        ) : (
          <Text>{data.username}</Text>
        )}
      </Box>
    </Flex>
  );
};
