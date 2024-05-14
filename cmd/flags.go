package cmd

import (
    "fmt"
    "github.com/integrii/flaggy"
)

var (
    dateFlag  string
    aboutFlag = false
)

func ParseFlags() {
    flaggy.DefaultParser.ShowVersionWithVersionFlag = false
    flaggy.DefaultParser.ShowHelpOnUnexpected = true

    flaggy.SetName("imir")
    flaggy.SetDescription("Is Mercury In Retrograde?")

    flaggy.String(&dateFlag, "d", "date", "Check if Mercury was in retrograde during a specific date.")
    flaggy.Bool(&aboutFlag, "a", "about", "Show information on the concept behind Mercury Retrograde.")

    flaggy.Parse()

	// subcommand?       
		fmt.Println("--date in use")
	}

	if aboutFlag {
		fmt.Println("--config in use")
	}

}