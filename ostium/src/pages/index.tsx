import { Channel } from "@liveassist/liber";
import { GetServerSideProps } from "next";
import { PageLayout } from "../components/PageLayout";
import { api } from "../lib/api";
import { Channels } from "../ui/channels/Channels";
import Head from "next/head";
import { DefaultHead } from "../components/DefaultHead";

interface IndexProps {
  channels: Channel[];
}

const Index: React.FC<IndexProps> = ({ channels }) => {
  return (
    <PageLayout withHead>
      <Channels channels={channels} />
    </PageLayout>
  );
};

export const getServerSideProps: GetServerSideProps = async (_context) => {
  const channels = await api.channels.list();
  return {
    props: {
      channels,
    },
  };
};

export default Index;
