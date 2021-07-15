import React from "react";
import { AuthPage } from "../components/AuthPage";
import { PageLayout } from "../components/PageLayout";
import { CreateChannel } from "../ui/channels/CreateChannel";

export interface CreateChannelPageProps {}

const CreateChannelPage: React.FC<CreateChannelPageProps> = ({}) => {
  return (
    <PageLayout withHead>
      <AuthPage>
        <CreateChannel />
      </AuthPage>
    </PageLayout>
  );
};

export default CreateChannelPage;
