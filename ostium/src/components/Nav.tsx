import { Flex } from "@chakra-ui/layout";
import { Box, Heading, Link } from "@chakra-ui/react";
import NextLink from "next/link";
import React from "react";

export interface NavProps {}

export const Nav: React.FC<NavProps> = ({}) => {
  return (
    <Flex top={0} p={4}>
      <Box ml={4}>
        <NextLink href="/">
          <Link
            fontSize="3xl"
            fontWeight="bold"
            bgGradient="linear(to-l, #7928CA,#FF0080)"
            bgClip="text"
          >
            LiveAssist
          </Link>
        </NextLink>
      </Box>
    </Flex>
  );
};