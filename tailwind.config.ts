import { Config } from "tailwindcss";
import forms from "@tailwindcss/forms";
import typography from "@tailwindcss/typography";
import templewind from "./pkg/plugin";
import themes from "./pkg/themes";

const themeColors = themes.themes.light; // Default to light theme

const config: Config = {
  content: ["./pkg/**/*.{js,ts,scss}"],
  theme: {
    extend: {
      colors: themeColors,
    },
  },
  plugins: [forms, typography, templewind],
};

export default config;
