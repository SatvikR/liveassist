import {
  Box,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  HStack,
  Input,
  Tag,
  TagCloseButton,
  TagLabel,
} from "@chakra-ui/react";
import { Form, Formik, FormikProps } from "formik";
import { useRouter } from "next/router";
import React, { useState } from "react";
import { InputField } from "../../components/InputField";
import { StyledButton } from "../../components/StyledButton";
import { useCreateChannel } from "../../lib/api-hooks/useCreateChannel";

export interface CreateChannelProps {}

interface CreateChannelValues {
  title: string;
}

export const CreateChannel: React.FC<CreateChannelProps> = ({}) => {
  const [keywords, setKeywords] = useState<string[]>([]);
  const [kwInput, setkwInput] = useState<string>("");
  const router = useRouter();
  const createChannel = useCreateChannel();

  const addKeyword = () => {
    if (kwInput != "") {
      setKeywords([...keywords, kwInput]);
      setkwInput("");
    }
  };

  const removeKeyword = (key: number) => {
    setKeywords(keywords.filter((_, i) => i !== key));
  };

  return (
    <>
      <Heading mb={4}>Create Channel</Heading>
      <Formik
        initialValues={{ title: "" }}
        onSubmit={async ({ title }, { setSubmitting }) => {
          await createChannel(title, keywords);
          router.push("/");
        }}
      >
        {(props: FormikProps<CreateChannelValues>) => (
          <Form>
            <InputField
              name="title"
              label="Title"
              placeholder="How do I reverse a linked list in Java?"
            />
            <Box my={4}>
              <FormControl>
                <FormLabel>Keywords</FormLabel>
                <Flex>
                  <Input
                    placeholder="ex. rest-api"
                    maxW="300px"
                    mr={4}
                    value={kwInput}
                    onChange={(e) => setkwInput(e.target.value)}
                  />
                  <StyledButton onClick={addKeyword}>Add</StyledButton>
                </Flex>
                <HStack my={4}>
                  {keywords.map((e, i) => (
                    <Tag
                      size="lg"
                      key={i}
                      px={4}
                      pb={0.5}
                      variant="solid"
                      colorScheme="purple"
                      borderRadius="full"
                    >
                      <TagLabel>{e}</TagLabel>
                      <TagCloseButton mt={1} onClick={() => removeKeyword(i)} />
                    </Tag>
                  ))}
                </HStack>
              </FormControl>
            </Box>
            <StyledButton type="submit">Submit</StyledButton>
          </Form>
        )}
      </Formik>
    </>
  );
};
