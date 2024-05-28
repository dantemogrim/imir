package messages

import (
	"fmt"
	"time"

	"github.com/dantemogrim/imir/pkg/styles"
)

type message struct {
	IMIR bool
	text string
}

var (
	retrograde    = message{true, "👹\n\nMercury %s in retrograde"}
	notRetrograde = message{false, "🕊️\n\nMercury %s *not* in retrograde"}
)

func Result(IMIR bool) string {
	if !IMIR {
		return styles.NotRetrograde(fmt.Sprintf(notRetrograde.text, "is") + ".")
	}
	return styles.Retrograde(fmt.Sprintf(retrograde.text, "is") + ".")
}

func DatedResult(givenDate time.Time, IMIR bool) string {
	var tense string = tenseByDate(givenDate)
	if !IMIR {
		return styles.NotRetrograde(fmt.Sprintf(notRetrograde.text, tense) + " " + givenDate.Format("2006-01-02") + ".")
	}
	return styles.Retrograde(fmt.Sprintf(retrograde.text, tense) + " " + givenDate.Format("2006-01-02") + ".")
}

func tenseByDate(givenDate time.Time) string {
	var today time.Time = time.Now().UTC().Truncate(24 * time.Hour)
	var response string

	switch {
	case givenDate.Before(today):
		response = "was"
	case givenDate.After(today):
		response = "will be"
	default:
		response = "is"
	}

	return response
}
