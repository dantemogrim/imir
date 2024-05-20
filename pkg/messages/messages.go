package messages

import (
	"time"
)

const ConnFailure string = "Wasn't able to connect to the Mercury Retrograde API. Status: "

func DatedResult(givenDate time.Time, IMIR bool) string {
    tense := datedTense(givenDate)
    if !IMIR {
        return "Mercury " + tense[1] + " in retrograde " + givenDate.Format("2006-01-02") + "."
    }
    return "Mercury " + tense[0] + " in retrograde " + givenDate.Format("2006-01-02") + "."
}

func datedTense(givenDate time.Time) []string {
    today := time.Now()
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
