package messages

import (
	"testing"
	"time"
)

func TestResult(t *testing.T) {
	if Result(true) != "Mercury is in retrograde." {
		t.Errorf("Result(true) failed, expected 'Mercury is in retrograde.', got %v", Result(true))
	}

	if Result(false) != "Mercury is not in retrograde." {
		t.Errorf("Result(false) failed, expected 'Mercury is not in retrograde.', got %v", Result(false))
	}
}

func TestDatedResult(t *testing.T) {
	var past time.Time = time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC)
	var future time.Time = time.Now().AddDate(1, 0, 0)

	if DatedResult(past, true) != "Mercury was in retrograde 2021-12-31." {
		t.Errorf("DatedResult(past, true) failed, expected 'Mercury was in retrograde 2021-12-31.', got %v", DatedResult(past, true))
	}

	if DatedResult(future, false) != "Mercury won't be in retrograde "+future.Format("2006-01-02")+"." {
		t.Errorf("DatedResult(future, false) failed, expected 'Mercury won't be in retrograde %v.', got %v", future.Format("2006-01-02"), DatedResult(future, false))
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
