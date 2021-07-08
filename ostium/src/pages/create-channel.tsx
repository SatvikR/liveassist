import { Heading } from "@chakra-ui/react";
import React from "react";
import { AuthPage } from "../components/AuthPage";
import { PageLayout } from "../components/PageLayout";

export interface CreateChannelPageProps {}

const CreateChannelPage: React.FC<CreateChannelPageProps> = ({}) => {
  return (
    <AuthPage>
      <PageLayout>
        <Heading>Create Channel</Heading>
      </PageLayout>
    </AuthPage>
  );
};

export default CreateChannelPage;
