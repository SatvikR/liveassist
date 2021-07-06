import React from "react";
import { DefaultPage } from "../components/DefaultPage";
import { Login } from "../ui/users/Login";

export interface LoginPageProps {}

const LoginPage: React.FC<LoginPageProps> = ({}) => {
  return (
    <DefaultPage>
      <Login />
    </DefaultPage>
  );
};

export default LoginPage;
