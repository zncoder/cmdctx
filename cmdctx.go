package cmdctx

import (
	"flag"

	"github.com/zncoder/assert"
	"github.com/zncoder/ctx"
)

var verbose = flag.Bool("verbose", false, "enable verbose trace")

func New() ctx.Context {
	assert.OKf(flag.Parsed(), "New must be called after flag.Parse is called")
	cx := ctx.New(nil)
	if *verbose {
		cx = cx.WithLog()
	}
	return cx
}
