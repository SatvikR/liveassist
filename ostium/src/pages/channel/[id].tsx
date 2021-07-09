import { Heading } from "@chakra-ui/react";
import { useRouter } from "next/dist/client/router";
import React from "react";
import { AuthPage } from "../../components/AuthPage";
import { PageLayout } from "../../components/PageLayout";

export interface MessagingPageProps {}

const MessagingPage: React.FC<MessagingPageProps> = ({}) => {
  const router = useRouter();
  const { id } = router.query;

  return (
    <AuthPage>
      <PageLayout>
        <Heading>{id}</Heading>
      </PageLayout>
    </AuthPage>
  );
};

export default MessagingPage;
