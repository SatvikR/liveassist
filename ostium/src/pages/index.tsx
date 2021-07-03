import { Container } from "../components/Container";
import { Nav } from "../components/Nav";
import { Channels } from "../modules/channels/Channels";

const Index = () => {
  return (
    <>
      <Nav />
      <Container>
        <Channels />
      </Container>
    </>
  );
};

export default Index;
