import plugin from "tailwindcss/plugin";
import themes, { Themes } from "./themes";

function generateThemeStyles(themeNames: Array<keyof Themes>, themes: Themes) {
  const themeStyles: Record<string, Record<string, string>> = {};

  themeNames.forEach((themeName) => {
    const themeConfig = themes[themeName];
    const cssVariables = Object.entries(themeConfig).reduce(
      (acc: Record<string, string>, [key, value]) => {
        acc[`--${key}`] = value;
        return acc;
      },
      {}
    );

    themeStyles[`[data-theme="${themeName}"]`] = cssVariables;
  });

  return themeStyles;
}

function generateComponents(themeNames: Array<keyof Themes>, themes: Themes) {
  const components: Record<string, any> = {
    ".rounded-box": {
      borderRadius: "0.5rem",
    },
  };

  themeNames.forEach((themeName) => {
    const themeConfig = themes[themeName];

    Object.keys(themeConfig).forEach((colorKey) => {
      components[`.text-${colorKey}-content`] = {
        color: `var(--${colorKey})`,
      };
      components[`.border-${colorKey}-20`] = {
        borderColor: `var(--${colorKey})`,
        borderOpacity: 0.2,
      };
      components[`.bg-${colorKey}`] = {
        backgroundColor: `var(--${colorKey})`,
      };
      components[`.text-${colorKey}`] = {
        color: `var(--${colorKey})`,
      };
    });
  });

  return components;
}

function generateUtilities(themeNames: Array<keyof Themes>, themes: Themes) {
  const utilities: Record<string, any> = {};

  themeNames.forEach((themeName) => {
    const themeConfig = themes[themeName];

    Object.keys(themeConfig).forEach((colorKey) => {
      utilities[`.border-${colorKey}-200`] = {
        borderColor: `var(--${colorKey})`,
      };
    });
  });

  return utilities;
}

export default plugin(function ({ addBase, addComponents, addUtilities }) {
  const themeNames = Object.keys(themes.themes) as Array<keyof Themes>;

  addBase(generateThemeStyles(themeNames, themes.themes));
  addComponents(generateComponents(themeNames, themes.themes));
  addUtilities(generateUtilities(themeNames, themes.themes));
});
