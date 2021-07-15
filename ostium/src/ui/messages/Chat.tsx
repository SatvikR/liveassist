import { Box, Flex, Heading, Spinner, Stack } from "@chakra-ui/react";
import { Channel, Message as IMessage } from "@liveassist/liber";
import { Form, Formik, FormikProps } from "formik";
import React, { useEffect, useRef, useState } from "react";
import { InputField } from "../../components/InputField";
import { StyledButton } from "../../components/StyledButton";
import { useChannel } from "../../lib/api-hooks/useChannel";
import { useMessageClient } from "../../lib/api-hooks/useMessageClient";
import { Message } from "./Message";

export interface ChatProps {
  id: string;
  channel: Channel;
}

interface MessageForm {
  message: string;
}

export const Chat: React.FC<ChatProps> = ({ id, channel }) => {
  const [messages, setMessages] = useState<IMessage[]>([]);
  const bottomMsg = useRef<HTMLDivElement>(null);

  const { isLoading, data, isError } = useChannel(id, channel);
  const { isConnecting, client } = useMessageClient(
    id,
    (m) => {
      setMessages((oldMessages) => [m, ...oldMessages]);
      bottomMsg.current.scrollIntoView();
    },
    (m) => {
      setMessages(m);
      if (bottomMsg.current) {
        bottomMsg.current.scrollIntoView();
      }
    }
  );

  useEffect(() => {
    if (bottomMsg.current) {
      bottomMsg.current.scrollIntoView();
    }

    return;
  }, [bottomMsg.current]);

  useEffect(() => {
    return () => {
      client.current.close();
    };
  }, []);

  if (isLoading || isError) {
    return (
      <Box>
        <Spinner
          thickness="4px"
          speed="0.65s"
          emptyColor="gray.200"
          color="blue.500"
          size="xl"
        />
      </Box>
    );
  }

  return (
    <>
      {isConnecting ? (
        <Box>
          <Heading>Connecting...</Heading>
        </Box>
      ) : (
        <Box>
          <Box mb={4}>
            <Heading>{data.name}</Heading>
          </Box>
          <Stack
            overflowY="auto"
            height="75vh"
            spacing={4}
            flexDir="column-reverse"
          >
            <Box ref={bottomMsg}></Box>
            {messages.map((e, i) => (
              <Message isFirst={i == 0} message={e} key={i} />
            ))}
          </Stack>
          <Formik
            initialValues={{ message: "" }}
            onSubmit={async ({ message }, { setSubmitting, resetForm }) => {
              resetForm();
              client.current.send(message);
              setSubmitting(false);
            }}
          >
            {(props: FormikProps<MessageForm>) => (
              <Form>
                <Flex>
                  <InputField
                    name="message"
                    label=""
                    placeholder="Send a message"
                  />
                  <StyledButton
                    ml={4}
                    my="auto"
                    type="submit"
                    isLoading={props.isSubmitting}
                  >
                    Send
                  </StyledButton>
                </Flex>
              </Form>
            )}
          </Formik>
        </Box>
      )}
    </>
  );
};
