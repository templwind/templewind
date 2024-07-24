import { defineConfig } from "vite";
import { resolve } from "path";
import { readdirSync } from "fs";
import dts from "vite-plugin-dts";
import { libInjectCss } from "vite-plugin-lib-inject-css";

// Function to get all components from the pkg/components directory
function getComponents() {
  const componentsDir = resolve(__dirname, "pkg/components");
  const components = readdirSync(componentsDir);

  const entry: Record<string, string> = {};
  const input: Record<string, string> = {};

  components.forEach((component) => {
    const componentName = component.replace(/\.[^/.]+$/, ""); // Remove file extension
    const componentPath = resolve(componentsDir, component);

    entry[componentName] = componentPath;
    input[componentName] = componentPath;
  });

  return { entry, input };
}

const { entry, input } = getComponents();

export default defineConfig({
  plugins: [
    libInjectCss(),
    dts({
      include: ["pkg/**/*.ts"],
      outDir: "dist/types",
      insertTypesEntry: true,
    }),
  ],
  build: {
    copyPublicDir: false,
    lib: {
      entry: {
        main: resolve(__dirname, "pkg/main.ts"),
        ...entry,
      },
    },
    rollupOptions: {
      external: [],
      input: {
        main: resolve(__dirname, "pkg/main.ts"),
        ...input,
      },
      output: [
        {
          format: "es",
          entryFileNames: (chunk) => {
            if (chunk.name === "main") {
              return "index.esm.js"; // ESM entry output directly in the dist folder
            } else {
              return "components/[name].esm.js"; // Component outputs in the dist/components folder
            }
          },
          assetFileNames: "assets/[name][extname]",
          dir: resolve(__dirname, "dist"),
        },
        {
          format: "cjs",
          entryFileNames: (chunk) => {
            if (chunk.name === "main") {
              return "index.cjs.js"; // CommonJS entry output directly in the dist folder
            } else {
              return "components/[name].cjs.js"; // Component outputs in the dist/components folder
            }
          },
          assetFileNames: "assets/[name][extname]",
          dir: resolve(__dirname, "dist"),
        },
      ],
    },
  },
});
