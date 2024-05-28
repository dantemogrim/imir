package flags

import (
	"flag"
)

var (
	about *bool   = flag.Bool("a", false, "Learn more about Mercury in Retrograde.")
	date  *string = flag.String("d", "1991-12-31", "Was Mercury in retrograde on a specific date?")
	help  *bool   = flag.Bool("h", false, "Print help message.")
)

func Init() {
	flag.Parse()

	if *help {
		helpOption()
	}

	if *date != "" {
		dateOption()
	}

	if *about {
		aboutOption()
	}
}
