import { ChakraProvider } from "@chakra-ui/react";
import { AppProps } from "next/app";
import { QueryClientProvider } from "react-query";
import { queryClient } from "../lib/queryClient";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <ChakraProvider resetCSS>
        <Component {...pageProps} />
      </ChakraProvider>
    </QueryClientProvider>
  );
}

export default MyApp;
