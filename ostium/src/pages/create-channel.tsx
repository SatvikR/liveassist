import React from "react";
import { AuthPage } from "../components/AuthPage";
import { PageLayout } from "../components/PageLayout";
import { CreateChannel } from "../ui/channels/CreateChannel";

export interface CreateChannelPageProps {}

const CreateChannelPage: React.FC<CreateChannelPageProps> = ({}) => {
  return (
    <AuthPage>
      <PageLayout>
        <CreateChannel />
      </PageLayout>
    </AuthPage>
  );
};

export default CreateChannelPage;
