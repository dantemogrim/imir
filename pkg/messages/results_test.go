package messages

import (
	"fmt"
	"testing"
	"time"

	"github.com/dantemogrim/imir/pkg/styles"
)

func TestResult(t *testing.T) {
	var expectedRetrograde string = styles.Retrograde(retrogradeTxt)
	var expectedNotRetrograde string = styles.NotRetrograde(notRetrogradeTxt)

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

	var expectedRetrograde string = styles.Retrograde(fmt.Sprintf(datedRetrogradeTxt, "was", past.Format("2006-01-02")))
	var expectedNotRetrograde string = styles.NotRetrograde(fmt.Sprintf(datedNotRetrogradeTxt, "won't be", future.Format("2006-01-02")))

	if DatedResult(past, true) != expectedRetrograde {
		t.Errorf("DatedResult(past, true) failed, expected '"+expectedRetrograde+"', got %v", DatedResult(past, true))
	}

	if DatedResult(future, false) != expectedNotRetrograde {
		t.Errorf("DatedResult(future, false) failed, expected '"+expectedNotRetrograde+"' got %v", DatedResult(future, false))
	}
}

func TestDatedTense(t *testing.T) {
	past := time.Now().AddDate(0, 0, -1)
	future := time.Now().AddDate(0, 0, 1)
	now := time.Now().UTC().Truncate(24 * time.Hour)

	if datedTense(past)[0] != "was" || datedTense(past)[1] != "wasn't" {
		t.Errorf("datedTense(past) failed, expected ['was', 'wasn't'], got %v", datedTense(past))
	}

	if datedTense(future)[0] != "will be" || datedTense(future)[1] != "won't be" {
		t.Errorf("datedTense(future) failed, expected ['will be', 'won't be'], got %v", datedTense(future))
	}

	if datedTense(now)[0] != "is" || datedTense(now)[1] != "is not" {
		t.Errorf("datedTense(now) failed, expected ['is', 'is not'], got %v", datedTense(now))
	}
}
