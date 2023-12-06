
import React from "react";
import type { AppProps } from "next/app";
import { QueryClient, QueryClientProvider, hydrate } from "@tanstack/react-query";
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import RootLayout, { Layout } from "@/app/layout";
import "../app/layout"
import { UserProvider } from "@auth0/nextjs-auth0/client";


function MyApp({ Component, pageProps }: AppProps) {
  const [queryClient] = React.useState(() => new QueryClient());
  return (
    <UserProvider>
      <Layout>
        <QueryClientProvider client={queryClient}>
          {/* <Hydrate state={pageProps.dehydratedState}> */}
          <Component {...pageProps} />
          {/* </Hydrate> */}
          <ReactQueryDevtools initialIsOpen={false} position="left" />
        </QueryClientProvider>
      </Layout>
    </UserProvider>
  );
}

export default MyApp;