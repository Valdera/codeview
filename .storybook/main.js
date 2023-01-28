// @ts-check
const path = require('path');

module.exports = {
  stories: ['../**/*.stories.mdx', '../**/*.stories.@(js|jsx|ts|tsx)'],
  staticDirs: ['../public'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',
    '@storybook/addon-interactions',
    'storybook-addon-next',
    'storybook-css-modules-preset',
    'storybook-addon-next-router',
    '@chakra-ui/storybook-addon',
    {
      /**
       * Fix Storybook issue with PostCSS@8
       * @see https://github.com/storybookjs/storybook/issues/12668#issuecomment-773958085
       */
      name: '@storybook/addon-postcss',
      options: {
        postcssLoaderOptions: {
          implementation: require('postcss'),
        },
      },
    },
  ],
  features: {
    emotionAlias: false,
  },
  framework: '@storybook/react',
  core: {
    builder: '@storybook/builder-webpack5',
  },
  webpackFinal: async (config) => {
    // Return the altered config
    return {
      ...config,
      resolve: {
        ...config.resolve,
        // can likely be removed in storybook > 6.4 see - https://github.com/storybookjs/storybook/issues/17458
        fallback: {
          ...config.resolve.fallback,
          assert: require.resolve('assert'),
        },
      },
    };
  },
};
