package messages

import (
	"time"
)

func Result(IMIR bool) string {
	if !IMIR {
		// TODO styles
		return "Mercury is not in retrograde."
	}
	// TODO retrograde styles (see bubbletea examples)
	return "Mercury is in retrograde."
}

func DatedResult(givenDate time.Time, IMIR bool) string {
	var tense []string = datedTense(givenDate)
	if !IMIR {
		return "Mercury " + tense[1] + " in retrograde " + givenDate.Format("2006-01-02") + "."
	}
	return "Mercury " + tense[0] + " in retrograde " + givenDate.Format("2006-01-02") + "."
}

func datedTense(givenDate time.Time) []string {
	var today time.Time = time.Now().UTC().Truncate(24 * time.Hour)
	var trueResponse, falseResponse string

	switch {
	case givenDate.Before(today):
		trueResponse = "was"
		falseResponse = "wasn't"
	case givenDate.After(today):
		trueResponse = "will be"
		falseResponse = "won't be"
	default:
		trueResponse = "is"
		falseResponse = "is not"
	}

	return []string{trueResponse, falseResponse}
}
