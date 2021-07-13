import { useRouter } from "next/router";
import React from "react";
import { useLoggedIn } from "../lib/state/useLoggedIn";
import { Nav } from "./Nav";

export interface AuthPageProps {}

export const AuthPage: React.FC<AuthPageProps> = ({ children }) => {
  const loginLoading = useLoggedIn((state) => state.loading);
  const loggedIn = useLoggedIn((state) => state.loggedIn);
  const router = useRouter();

  if (loginLoading) {
    return (
      <>
        <Nav />
      </>
    );
  }

  if (!loginLoading && loggedIn) {
    return <>{children}</>;
  }
  router.push("/login");
  return <></>;
};
