import tailwindcss from "tailwindcss";
// import autoprefixer from "autoprefixer";
import postcssImport from "postcss-import";
import nesting from "tailwindcss/nesting/index.js"; // Updated import

const config = {
  plugins: [postcssImport, nesting, tailwindcss],
};

export default config;
