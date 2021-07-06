import { Box, Heading } from "@chakra-ui/react";
import { Form, Formik, FormikProps } from "formik";
import { useRouter } from "next/dist/client/router";
import React from "react";
import { Container } from "../../components/Container";
import { InputField } from "../../components/InputField";
import { StyledButton } from "../../components/StyledButton";
import { useSignup } from "../../lib/api-hooks/useSignup";

export interface SignupProps {}

interface SignupValues {
  username: string;
  email: string;
  password: string;
}

export const Signup: React.FC<SignupProps> = ({}) => {
  const signup = useSignup();
  const router = useRouter();

  return (
    <Container size="small">
      <Heading>Signup</Heading>
      <Formik
        initialValues={{ username: "", password: "", email: "" }}
        onSubmit={async (
          { username, email, password },
          { setSubmitting, setErrors }
        ) => {
          const errors = await signup(username, email, password);
          if (!errors) {
            router.push("/");
            return;
          }

          setErrors(errors);
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
