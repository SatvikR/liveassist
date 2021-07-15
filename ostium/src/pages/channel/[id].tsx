import { Channel } from "@liveassist/liber";
import { GetServerSideProps } from "next";
import React from "react";
import { AuthPage } from "../../components/AuthPage";
import { DefaultHead } from "../../components/DefaultHead";
import { PageLayout } from "../../components/PageLayout";
import { api } from "../../lib/api";
import { Chat } from "../../ui/messages/Chat";

export interface MessagingPageProps {
  channel: Channel;
}

const MessagingPage: React.FC<MessagingPageProps> = ({ channel }) => {
  const meta = {
    title: `${channel.name} | LiveAssist`,
    description: `Keywords: ${channel.keywords.join(", ")}. ${channel.name}`,
    url: `https://liveassist.satvikreddy.com/channel/${channel.id}`,
  };

  return (
    <PageLayout>
      <DefaultHead {...meta} />
      <AuthPage>
        <Chat id={channel.id as string} channel={channel} />
      </AuthPage>
    </PageLayout>
  );
};

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { id } = context.query;
  const channel = await api.channels.get(id as string);

  return {
    props: {
      channel,
    },
  };
};

export default MessagingPage;
