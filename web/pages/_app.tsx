
import React from "react";
import type { AppProps } from "next/app";
import { Hydrate, QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import RootLayout, { Layout } from "@/app/layout";
import "../app/layout"


function MyApp({ Component, pageProps }: AppProps) {
  const [queryClient] = React.useState(() => new QueryClient());
  return (
    <Layout>
      <QueryClientProvider client={queryClient}>
      <Hydrate state={pageProps.dehydratedState}>
          <Component {...pageProps} />
        <ReactQueryDevtools initialIsOpen={false} />
</Hydrate>
      </QueryClientProvider> 
     </Layout>
  );
}

export default MyApp;