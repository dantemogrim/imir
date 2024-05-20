package utils

import (
	"os"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"golang.org/x/term"
)

const ClearScreen string = "\x1bc"

func SpinnerStart() {
	spinner := wow.New(os.Stdout, spin.Get(spin.Star), " Stargazing.. 🔭")
	spinner.Start()
}

func TerminalDimensions() ([]int, error) {
	width, height, err := term.GetSize(0)
	if err != nil {
		return nil, err
	}
	return []int{width, height}, nil
}
