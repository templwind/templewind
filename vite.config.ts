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
  plugins: [libInjectCss(), dts({ include: ["lib"] })],
  build: {
    copyPublicDir: false,
    lib: {
      entry: {
        main: resolve(__dirname, "pkg/main.ts"),
        ...entry,
      },
      formats: ["es"],
    },
    rollupOptions: {
      external: [],
      input: {
        main: resolve(__dirname, "pkg/main.ts"),
        ...input,
      },
      output: {
        format: "es",
        assetFileNames: "assets/[name][extname]",
        entryFileNames: (chunk) => {
          // Check the chunk name to determine the output directory
          if (chunk.name === "main") {
            return "index.js"; // Main entry output directly in the dist folder
          } else {
            return "components/[name].js"; // Component outputs in the dist/components folder
          }
        },
        dir: resolve(__dirname, "dist"),
      },
    },
  },
});
