package messages

import (
	"time"

	"github.com/dantemogrim/imir/pkg/styles"
)

func Result(IMIR bool) string {
	if !IMIR {
		return styles.NotRetrograde("🕊️\n\nMercury is *not* in retrograde.")
	}
	return styles.Retrograde("👹\n\nMercury is in retrograde.")
}

func DatedResult(givenDate time.Time, IMIR bool) string {
	var tense []string = datedTense(givenDate)
	if !IMIR {
		return styles.NotRetrograde("🕊️\n\nMercury " + tense[1] + " in retrograde " + givenDate.Format("2006-01-02") + ".")
	}
	return styles.Retrograde("👹\n\nMercury " + tense[0] + " in retrograde " + givenDate.Format("2006-01-02") + ".")
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
