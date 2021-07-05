import { Channel } from "@liveassist/liber";
import { GetServerSideProps } from "next";
import { DefaultPage } from "../components/DefaultPage";
import { api } from "../lib/api";
import { Channels } from "../modules/channels/Channels";

interface IndexProps {
  channels: Channel[];
}

const Index: React.FC<IndexProps> = ({ channels }) => {
  return (
    <DefaultPage>
      <Channels channels={channels} />
    </DefaultPage>
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
