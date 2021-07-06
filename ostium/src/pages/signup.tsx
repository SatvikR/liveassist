import React from "react";
import { DefaultPage } from "../components/DefaultPage";
import { Signup } from "../ui/users/Signup";

export interface SignupPageProps {}

const SignupPage: React.FC<SignupPageProps> = ({}) => {
  return (
    <DefaultPage>
      <Signup />
    </DefaultPage>
  );
};

export default SignupPage;
