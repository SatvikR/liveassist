import React from "react";
import { Container } from "./Container";
import { DefaultHead } from "./DefaultHead";
import { Nav } from "./Nav";

export interface DefaultPageProps {
  withHead?: boolean;
}

export const PageLayout: React.FC<DefaultPageProps> = ({
  withHead,
  children,
}) => {
  return (
    <>
      {withHead && <DefaultHead />}
      <Nav />
      <Container>{children}</Container>
    </>
  );
};
