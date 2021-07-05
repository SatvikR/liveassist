import React from "react";
import { Container } from "./Container";
import { Nav } from "./Nav";

export interface DefaultPageProps {}

export const DefaultPage: React.FC<DefaultPageProps> = ({ children }) => {
  return (
    <>
      <Nav />
      <Container>{children}</Container>
    </>
  );
};
