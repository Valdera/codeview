import { extendTheme } from '@chakra-ui/react';
const { getTheme } = require('./styles/utils/config');

const theme = getTheme('chakra');

export default extendTheme({
  fonts: theme.fontFamily(),
  colors: theme.colors(),
});
