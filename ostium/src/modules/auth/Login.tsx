import { Box, Button, Heading } from "@chakra-ui/react";
import { Form, Formik, FormikProps } from "formik";
import React from "react";
import { InputField } from "../../components/InputField";
import { api } from "../../lib/api";

export interface LoginProps {}
interface LoginValues {
  username: string;
  password: string;
}

export const Login: React.FC<LoginProps> = ({}) => {
  return (
    <>
      <Box maxW="800px" p={4} mx="auto">
        <Heading>Login</Heading>
        <Formik
          initialValues={{ username: "", password: "" }}
          onSubmit={({ username, password }, { setSubmitting }) => {
            api.users
              .login(username, password)
              .then((data) => console.log(data))
              .then(() => setSubmitting(false));
          }}
        >
          {(props: FormikProps<LoginValues>) => (
            <Form>
              <InputField
                name="username"
                label="username"
                placeholder="username"
              />
              <InputField
                name="password"
                label="password"
                placeholder="password"
                type="password"
              />
              <Button
                mt={4}
                colorScheme="teal"
                isLoading={props.isSubmitting}
                type="submit"
              >
                Submit
              </Button>
            </Form>
          )}
        </Formik>
      </Box>
    </>
  );
};
