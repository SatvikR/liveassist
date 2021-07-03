import { Box } from "@chakra-ui/react";
import React from "react";

export interface ContainerProps {}

export const Container: React.FC<ContainerProps> = ({ children }) => {
  return (
    <Box maxW="1200px" mx="auto" px={4}>
      {children}
    </Box>
  );
};
