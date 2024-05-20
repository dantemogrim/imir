package flags

import (
	"os"

	"github.com/dantemogrim/imir/pkg/pager"
)

func aboutOption () {
	pager.Render()
	os.Exit(0)
}
