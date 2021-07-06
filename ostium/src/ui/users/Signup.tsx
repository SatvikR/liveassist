import { Box, Heading } from "@chakra-ui/react";
import { Form, Formik, FormikProps } from "formik";
import React from "react";
import { Container } from "../../components/Container";
import { InputField } from "../../components/InputField";
import { StyledButton } from "../../components/StyledButton";

export interface SignupProps {}

interface SignupValues {
  username: string;
  email: string;
  password: string;
}

export const Signup: React.FC<SignupProps> = ({}) => {
  return (
    <Container size="small">
      <Heading>Signup</Heading>
      <Formik
        initialValues={{ username: "", password: "", email: "" }}
        onSubmit={({}, { setSubmitting }) => {
          setSubmitting(false);
        }}
      >
        {(props: FormikProps<SignupValues>) => (
          <Form>
            <InputField
              name="username"
              label="Username"
              placeholder="username"
            />
            <InputField
              name="email"
              label="Email"
              placeholder="email"
              type="email"
            />
            <InputField
              name="password"
              label="Password"
              placeholder="password"
              type="password"
            />
            <StyledButton mt={4} isLoading={props.isSubmitting} type="submit">
              Submit
            </StyledButton>
          </Form>
        )}
      </Formik>
    </Container>
  );
};
