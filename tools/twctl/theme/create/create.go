package create

import "github.com/templwind/templwind/tools/twctl/internal/cobrax"

var (
	Cmd = cobrax.NewCommand("create", cobrax.WithRunE(nil))
)
