
import React from "react";
import type { AppProps } from "next/app";
import { Hydrate, QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import RootLayout, { Layout } from "@/app/layout";
import "../app/layout"
import { UserProvider } from "@auth0/nextjs-auth0/client";


function MyApp({ Component, pageProps }: AppProps) {
  const [queryClient] = React.useState(() => new QueryClient());
  return (
    <Layout>
      <UserProvider>
      <QueryClientProvider client={queryClient}>
          <Component {...pageProps} />
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider> 
      </UserProvider>
     </Layout>
  );
}

export default MyApp;