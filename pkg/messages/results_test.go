package messages

import (
	"fmt"
	"testing"
	"time"

	"github.com/dantemogrim/imir/pkg/styles"
)

func TestResult(t *testing.T) {
	var expectedRetrograde string = styles.Retrograde(fmt.Sprintf(retrograde.text, "is") + ".")
	var expectedNotRetrograde string = styles.NotRetrograde(fmt.Sprintf(notRetrograde.text, "is") + ".")

	if Result(true) != expectedRetrograde {
		t.Errorf("Result(true) failed, expected 'Mercury is in retrograde.', got %v", Result(true))
	}

	if Result(false) != expectedNotRetrograde {
		t.Errorf("Result(false) failed, expected 'Mercury is *not* in retrograde.', got %v", Result(false))
	}
}

func TestDatedResult(t *testing.T) {
	var past time.Time = time.Now().AddDate(-1, 0, 0)
	var future time.Time = time.Now().AddDate(1, 0, 0)

	var expectedRetrograde string = styles.Retrograde(fmt.Sprintf(retrograde.text, "was") + " " + past.Format("2006-01-02") + ".")
	var expectedNotRetrograde string = styles.NotRetrograde(fmt.Sprintf(notRetrograde.text, "will") + " " + future.Format("2006-01-02") + ".")

	if DatedResult(past, true) != expectedRetrograde {
		t.Errorf("DatedResult(past, true) failed, expected '"+expectedRetrograde+"', got %v", DatedResult(past, true))
	}

	if DatedResult(future, false) != expectedNotRetrograde {
		t.Errorf("DatedResult(future, false) failed, expected '"+expectedNotRetrograde+"' got %v", DatedResult(future, false))
	}
}

func TestTenseByDate(t *testing.T) {
	past := time.Now().AddDate(0, 0, -1)
	future := time.Now().AddDate(0, 0, 1)
	now := time.Now().UTC().Truncate(24 * time.Hour)

	if tenseByDate(past) != "was" {
		t.Errorf("tenseByDate(past) failed, expected ['was', 'wasn't'], got %v", tenseByDate(past))
	}

	if tenseByDate(future) != "will" {
		t.Errorf("tenseByDate(future) failed, expected ['will be', 'won't be'], got %v", tenseByDate(future))
	}

	if tenseByDate(now) != "is" {
		t.Errorf("tenseByDate(now) failed, expected ['is', 'is not'], got %v", tenseByDate(now))
	}
}
