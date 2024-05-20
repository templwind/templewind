package install

import "github.com/templwind/templwind/tools/twctl/internal/cobrax"

var (
	Cmd = cobrax.NewCommand("install", cobrax.WithRunE(nil))
)

func init() {

}
