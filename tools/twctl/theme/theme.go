package theme

import (
	"github.com/templwind/templwind/tools/twctl/internal/cobrax"
	"github.com/templwind/templwind/tools/twctl/theme/create"
	"github.com/templwind/templwind/tools/twctl/theme/install"
)

var (
	Cmd = cobrax.NewCommand("theme", cobrax.WithRunE(nil))
)

func init() {
	Cmd.AddCommand(install.Cmd, create.Cmd)
}
