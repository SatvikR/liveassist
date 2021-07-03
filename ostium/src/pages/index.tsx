import { GetServerSideProps } from "next";
import { Container } from "../components/Container";
import { Nav } from "../components/Nav";
import { Channels } from "../modules/channels/Channels";
import { Channel } from "@liveassist/liber";
import { api } from "../lib/api";

interface IndexProps {
  channels: Channel[];
}

const Index: React.FC<IndexProps> = ({ channels }) => {
  return (
    <>
      <Nav />
      <Container>
        <Channels channels={channels} />
      </Container>
    </>
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
