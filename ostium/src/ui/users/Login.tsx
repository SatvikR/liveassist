import { Box, Heading, Link, Text } from "@chakra-ui/react";
import { Form, Formik, FormikProps } from "formik";
import { useRouter } from "next/dist/client/router";
import React from "react";
import { Container } from "../../components/Container";
import { InputField } from "../../components/InputField";
import { StyledButton } from "../../components/StyledButton";
import { useLogin } from "../../lib/api-hooks/useLogin";
import NextLink from "next/link";
import { gradient } from "../../components/constants";

export interface LoginProps {}
interface LoginValues {
  username: string;
  password: string;
}

export const Login: React.FC<LoginProps> = ({}) => {
  const login = useLogin();
  const router = useRouter();

  return (
    <Container size="small">
      <Heading>Login</Heading>
      <Formik
        initialValues={{ username: "", password: "" }}
        onSubmit={async (
          { username, password },
          { setSubmitting, setErrors }
        ) => {
          const errors = await login(username, password);
          if (!errors) {
            router.push("/");
            return;
          }
          setErrors(errors);
          setSubmitting(false);
        }}
      >
        {(props: FormikProps<LoginValues>) => (
          <Form>
            <InputField
              name="username"
              label="Username"
              placeholder="username"
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
      <Box mt={4}>
        <Text>
          Don't have an account?{" "}
          <NextLink href="/signup">
            <Link
              bgGradient={gradient}
              bgClip="text"
              _hover={{
                textDecoration: "gray",
                textDecorationLine: "underline",
              }}
            >
              Signup here
            </Link>
          </NextLink>
        </Text>
      </Box>
    </Container>
  );
};
