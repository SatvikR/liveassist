import React from "react";
import { PageLayout } from "../components/PageLayout";
import { Login } from "../ui/users/Login";

export interface LoginPageProps {}

const LoginPage: React.FC<LoginPageProps> = ({}) => {
  return (
    <PageLayout withHead>
      <Login />
    </PageLayout>
  );
};

export default LoginPage;
