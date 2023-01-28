import { ChakraProvider } from '@chakra-ui/react';
import { MantineProvider } from '@mantine/core';
import '@styles/globals.scss';
import type { AppProps } from 'next/app';
import chakraTheme from '../chakra.config';
import mantineTheme from '../mantine.config';
import { NextPageWithLayout } from './page';

interface AppPropsWithLayout extends AppProps {
  Component: NextPageWithLayout;
}

function MyApp({ Component, pageProps }: AppPropsWithLayout) {
  const getLayout = Component.getLayout || ((page) => page);

  return (
    <>
      <MantineProvider theme={mantineTheme} withGlobalStyles withNormalizeCSS>
        <ChakraProvider theme={chakraTheme}>
          {getLayout(<Component {...pageProps} />)}
        </ChakraProvider>
      </MantineProvider>
    </>
  );
}

export default MyApp;
