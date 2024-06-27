import { defineConfig } from "vite";

export default defineConfig({
    server: {
        open: false,
        port: 3000
    },
    build: {
        outDir: "assets",
        sourcemap: true,
        rollupOptions: {
            input: {
                main: "src/main.ts",
                styles: "src/styles.scss"
            },
            output: {
                entryFileNames: "js/[name].js",
                assetFileNames: (asset) => {
                    if (asset.name?.endsWith('.css')) {
                        return "css/[name][extname]";
                    }
                    return "assets/[name][extname]";
                }
            }
        }
    }
});
