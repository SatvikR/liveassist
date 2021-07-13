import { Heading } from "@chakra-ui/react";
import { useRouter } from "next/router";
import React from "react";
import { AuthPage } from "../../components/AuthPage";
import { PageLayout } from "../../components/PageLayout";
import { Chat } from "../../ui/messages/Chat";

export interface MessagingPageProps {}

const MessagingPage: React.FC<MessagingPageProps> = ({}) => {
  const router = useRouter();
  const { id } = router.query;

  return (
    <AuthPage>
      <PageLayout>
        <Chat id={id as string} />
      </PageLayout>
    </AuthPage>
  );
};

export default MessagingPage;
