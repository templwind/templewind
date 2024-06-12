package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/templwind/templwind/tools/twctl/internal/xo"
)

func init() {
	rootCmd.AddCommand(XoCmd())
}

func XoCmd() *cobra.Command {
	var inputPath, outputPath, baseImportPath string
	var additionalIgnoreTypes []string

	var ignoreTypes = map[string]bool{
		"ErrInsertFailed": true,
		"Error":           true,
		"ErrUpdateFailed": true,
		"ErrUpsertFailed": true,
		// Add other types to ignore as needed
	}

	var cmd = &cobra.Command{
		Use:   "parsexo",
		Short: "Parse .xo.go files",
		Run: func(cmd *cobra.Command, args []string) {
			for _, typeName := range additionalIgnoreTypes {
				ignoreTypes[typeName] = true
			}

			absInputPath, err := filepath.Abs(inputPath)
			if err != nil {
				fmt.Printf("Error resolving absolute path of input directory: %v\n", err)
				return
			}

			absOutputPath, err := filepath.Abs(outputPath)
			if err != nil {
				fmt.Printf("Error resolving absolute path of output directory: %v\n", err)
				return
			}

			err = xo.ProcessFiles(absInputPath, absOutputPath, ignoreTypes, baseImportPath)
			if err != nil {
				fmt.Printf("Error processing files: %v\n", err)
			}
		},
	}

	cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input directory path")
	cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output directory path")
	cmd.Flags().StringSliceVarP(&additionalIgnoreTypes, "ignore-types", "t", []string{}, "Additional types to ignore")
	cmd.Flags().StringVarP(&baseImportPath, "import-path", "b", "", "Base import path for the generated files")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")
	cmd.MarkFlagRequired("import-path")

	return cmd
}
