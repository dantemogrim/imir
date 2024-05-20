package flags

import (
	"github.com/integrii/flaggy"
)

var (
	aboutFlag = false
	dateFlag  string
)

func Init() {
	// TODO - replace flaggy with flag library.
	// TODO - add help flag.
	flaggy.DefaultParser.ShowVersionWithVersionFlag = false
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.SetName("imir")
	flaggy.SetDescription("Is Mercury In Retrograde?\n")
	flaggy.DefaultParser.AdditionalHelpPrepend = "http://github.com/dantemogrim/imir"

	flaggy.Bool(&aboutFlag, "a", "about", "Learn more about Mercury in Retrograde.")
	flaggy.String(&dateFlag, "d", "date YYYY-MM-DD", "Check if Mercury was in retrograde during a specific date.")

	flaggy.Parse()

	if dateFlag != "" {
		dateOption()
	}

	if aboutFlag {
		aboutOption()
	}

}
