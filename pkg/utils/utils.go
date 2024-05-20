package utils

import (
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

const ClearScreen string = "\x1bc"

func SpinnerStart() {
	spinner := wow.New(os.Stdout, spin.Get(spin.Star), " Stargazing.. 🔭")
	spinner.Start()
}
