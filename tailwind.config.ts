import { Config } from "tailwindcss";
import forms from "@tailwindcss/forms";
import typography from "@tailwindcss/typography";
import templewind from "./lib/plugin";
import themes from "./lib/themes";

const themeColors = themes.themes.light; // Default to light theme

const config: Config = {
  content: ["./src/**/*.{js,ts,scss}"],
  theme: {
    extend: {
      colors: themeColors,
    },
  },
  plugins: [forms, typography, templewind],
};

export default config;
