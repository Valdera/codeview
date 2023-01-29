const theme = require('../global.theme');

function isDict(v) {
  return (
    typeof v === 'object' &&
    v !== null &&
    !(v instanceof Array) &&
    !(v instanceof Date)
  );
}

const AVAILABLE_FRAMEWORK = ['tailwind', 'chakra', 'mantine'];

class Theme {
  constructor(framework) {
    if (!AVAILABLE_FRAMEWORK.includes(framework)) {
      throw new Error('Framework not supported');
    }

    this.framework = framework;
    this.data = theme;
  }

  colors() {
    function normalizeMantine(colors) {
      const res = {};
      const items = Object.entries(colors);
      for (let i = 0; i < items.length; i++) {
        const key = items[i][0];
        const itemColors = items[i][1];
        if (isDict(itemColors)) {
          const temp = [];
          for (let color in itemColors) {
            temp.push(itemColors[color]);
          }
          res[key] = temp;
        }
      }
      return res;
    }

    switch (this.framework) {
      case 'tailwind':
        return this.data['colors'];
      case 'chakra':
        return this.data['colors'];
      case 'mantine':
        return normalizeMantine(this.data['colors']);
      default:
        break;
    }
  }

  fontFamily() {
    function normalizeChakra(fonts) {
      const res = {};

      for (let key in fonts) {
        const value = fonts[key];
        let temp = ``;

        for (let j = 0; j < value.length; j++) {
          temp += `'${value[j]}'`;

          if (j != value.length - 1) {
            temp += ', ';
          }
        }

        res[key] = temp;
      }

      return res;
    }

    function normalizeMantine(fonts) {
      return {
        fontFamily: `${fonts.body[0]}, ${fonts.body[1]}`,
        fontFamilyMonospace: `${fonts.mono[0]}, ${fonts.mono[1]}`,
        headings: {
          fontFamily: `${fonts.heading[0]}, ${fonts.heading[1]}`,
        },
      };
    }

    switch (this.framework) {
      case 'tailwind':
        return this.data['fontFamily'];
      case 'chakra':
        return normalizeChakra(this.data['fontFamily']);
      case 'mantine':
        return normalizeMantine(this.data['fontFamily']);
      default:
        break;
    }
  }
}

function getTheme(framework) {
  const theme = new Theme(framework);
  return theme;
}

module.exports = { getTheme };
