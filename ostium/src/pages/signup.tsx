import React from "react";
import { PageLayout } from "../components/PageLayout";
import { Signup } from "../ui/users/Signup";

export interface SignupPageProps {}

const SignupPage: React.FC<SignupPageProps> = ({}) => {
  return (
    <PageLayout>
      <Signup />
    </PageLayout>
  );
};

export default SignupPage;
