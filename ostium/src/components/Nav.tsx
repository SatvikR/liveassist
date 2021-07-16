import {
  ArrowBackIcon,
  ArrowForwardIcon,
  HamburgerIcon,
} from "@chakra-ui/icons";
import { Flex } from "@chakra-ui/layout";
import {
  Box,
  IconButton,
  Link,
  Menu,
  MenuButton,
  MenuDivider,
  MenuItem,
  MenuList,
  Text,
} from "@chakra-ui/react";
import NextLink from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { useLogout } from "../lib/api-hooks/useLogout";
import { useMeQuery } from "../lib/api-hooks/useMeQuery";
import { gradient } from "./constants";

export interface NavProps {}

export const Nav: React.FC<NavProps> = ({}) => {
  const { isLoading, data, isError } = useMeQuery();
  const logout = useLogout();
  const router = useRouter();

  const goToProfile = () => {
    router.push("/profile");
  };

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
        {!isLoading &&
          (isError || !data?.username ? (
            <NextLink href="/login">
              <Link mr={8} my="auto">
                Login <ArrowForwardIcon mb={0.5} />
              </Link>
            </NextLink>
          ) : (
            <Flex>
              <Text fontSize="large" my="auto">
                {data.username}
              </Text>
              <Menu>
                <MenuButton
                  as={IconButton}
                  aria-label="options"
                  icon={<HamburgerIcon />}
                  variant="outline"
                  borderRadius={5}
                  mx={4}
                ></MenuButton>
                <MenuList>
                  <MenuItem onClick={goToProfile}>Profile</MenuItem>
                  <MenuDivider />
                  <MenuItem icon={<ArrowBackIcon />} onClick={logout}>
                    Logout
                  </MenuItem>
                </MenuList>
              </Menu>
            </Flex>
          ))}
      </Box>
    </Flex>
  );
};
