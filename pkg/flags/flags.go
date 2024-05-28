package flags

import (
	"flag"
)

// TODO flag bugs due to definitions here.
var (
	about *bool   = flag.Bool("a", false, "Learn more about Mercury in Retrograde.")
	date  *string = flag.String("d", "1991-12-31", "Was Mercury in retrograde on a specific date?")
	help  *bool   = flag.Bool("h", false, "Print help message.")
)

func Init() {
	flag.Parse()

	if !isFlagPassed("date") && *date != "" {
		dateOption()
	}

	if *help {
		helpOption()
	}

	if *about {
		aboutOption()
	}
}

func isFlagPassed(name string) bool {
	var found bool = false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
