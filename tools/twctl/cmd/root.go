package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/templwind/templwind/tools/twctl/internal/cobrax"
	"github.com/templwind/templwind/tools/twctl/internal/version"
	"github.com/templwind/templwind/tools/twctl/theme"
)

var (
	rootCmd = cobrax.NewCommand("twcli")
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH)

	rootCmd.AddCommand(theme.Cmd)

}
