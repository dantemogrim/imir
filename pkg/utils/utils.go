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

// To ease testing, we'll create this interface that we can mock.
type TerminalSizer interface {
	GetSize(fileDescriptor int) (width, height int, err error)
}

// Create a new type that..
type TerminalSize struct{}

// .. attaches a GetSize method onto the TerminalSize type.
func (TerminalSize) GetSize(fileDescriptor int) (width, height int, err error) {
	return term.GetSize(fileDescriptor)
}

var Terminal TerminalSizer = TerminalSize{}

func TerminalDimensions() ([]int, error) {
	width, height, err := Terminal.GetSize(0)
	if err != nil {
		return nil, err
	}
	return []int{width, height}, nil
}
