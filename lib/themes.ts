export interface Theme {
  [key: string]: string;
}

export interface Themes {
  [themeName: string]: Theme;
}

const themes: { themes: Themes } = {
  themes: {
    light: {
      primary: "#3490dc",
      secondary: "#ffed4a",
      accent: "#e3342f",
      neutral: "#f3f4f6",
      "base-100": "#ffffff",
      "base-200": "#f0f4f8",
      "base-300": "#d1d5db",
      "base-content": "#1f2937",
      info: "#3490dc",
      success: "#38c172",
      warning: "#ffed4a",
      error: "#e3342f",
    },
    dark: {
      primary: "#1d4ed8",
      secondary: "#facc15",
      accent: "#e11d48",
      neutral: "#111827",
      "base-100": "#1c1c1e",
      "base-200": "#2d2d2d",
      "base-300": "#3f3f3f",
      "base-content": "#ffffff",
      info: "#3490dc",
      success: "#38c172",
      warning: "#ffed4a",
      error: "#e3342f",
    },
  },
};

export default themes;
