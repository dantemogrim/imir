package messages

import (
	"fmt"
	"time"

	"github.com/dantemogrim/imir/pkg/styles"
)

const (
	retrogradeTxt         string = "👹\n\nMercury is in retrograde."
	notRetrogradeTxt      string = "️🕊️\n\nMercury is *not* in retrograde."
	datedRetrogradeTxt    string = "👹\n\nMercury %s in retrograde %s."
	datedNotRetrogradeTxt string = "️🕊️\n\nMercury %s in retrograde %s."
)

func Result(IMIR bool) string {
	if !IMIR {
		return styles.NotRetrograde(notRetrogradeTxt)
	}
	return styles.Retrograde(retrogradeTxt)
}

func DatedResult(givenDate time.Time, IMIR bool) string {
	var tense []string = datedTense(givenDate)
	if !IMIR {
		return styles.NotRetrograde(fmt.Sprintf(datedNotRetrogradeTxt, tense[1], givenDate.Format("2006-01-02")))
	}
	return styles.Retrograde(fmt.Sprintf(datedRetrogradeTxt, tense[0], givenDate.Format("2006-01-02")))
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
