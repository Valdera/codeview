const { getTheme } = require('./styles/utils/config');
const theme = getTheme('mantine');

const mantineTheme = {
  colors: theme.colors(),
  ...theme.fontFamily(),
};

export default mantineTheme;
