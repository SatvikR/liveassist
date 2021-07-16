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

const LOAD_BOUNDARY = 10;

const scrollRef = (r: React.MutableRefObject<HTMLDivElement>) => {
  if (r.current) {
    r.current.scrollIntoView({ behavior: "smooth" });
  }
};

export const Chat: React.FC<ChatProps> = ({ id, channel }) => {
  const [messages, setMessages] = useState<IMessage[]>([]);
  // hasMore stores whether or not the client can load more messages
  const [hasMore, setHasMore] = useState<boolean>(false);

  const bottomMsg = useRef<HTMLDivElement>(null);

  const { isLoading, data, isError } = useChannel(id, channel);
  const { isConnecting, client } = useMessageClient(
    id,
    (message) => {
      setMessages((currentMessages) => [message, ...currentMessages]);
      scrollRef(bottomMsg);
    },
    (messages) => {
      setHasMore(messages.length >= client.current.PAGE_SIZE);
      setMessages((currentMessages) => [...currentMessages, ...messages]);
    }
  );

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

  const handleScroll: React.UIEventHandler<HTMLDivElement> = (e) => {
    // Use scrollHeight + scrollTop to compare rather than scrollHeight - scrollTop
    // because we are technically checking a scroll to the bottom rather than a scroll to the top,
    // since the flex direction is reversed.

    // Use LOAD_BOUNDARY as a range to make the infinite scroll more seemless
    if (
      Math.abs(
        e.currentTarget.scrollHeight +
          e.currentTarget.scrollTop -
          e.currentTarget.clientHeight
      ) <= LOAD_BOUNDARY
    ) {
      if (hasMore) {
        // Make sure client cannot send another loadMore request during the current one
        setHasMore(false);
        client.current.loadMore(messages[messages.length - 1]);
      }
    }
  };

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
            onScroll={handleScroll}
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
