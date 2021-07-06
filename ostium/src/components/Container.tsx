import { Box } from "@chakra-ui/react";
import React from "react";

export interface ContainerProps {
  size?: "small" | "normal";
}

export const Container: React.FC<ContainerProps> = ({ children, size }) => {
  return (
    <Box
      maxW={size ? (size == "normal" ? "1200px" : "800px") : "1200px"}
      mx="auto"
      px={4}
    >
      {children}
    </Box>
  );
};
