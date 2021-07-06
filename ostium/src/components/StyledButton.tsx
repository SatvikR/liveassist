import { Button, ButtonProps } from "@chakra-ui/react";
import React from "react";
import { gradient } from "./constants";

export type StyledButtonProps = ButtonProps;

export const StyledButton: React.FC<StyledButtonProps> = ({
  children,
  ...props
}) => {
  return (
    <Button
      bgGradient={gradient}
      textColor="white"
      _hover={{
        bgGradient: { gradient },
      }}
      {...props}
    >
      {children}
    </Button>
  );
};
