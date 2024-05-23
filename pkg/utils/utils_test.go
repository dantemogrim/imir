package utils

import (
	"testing"
)

type MockTerminal struct{}

func (MockTerminal) GetSize(fd int) (width, height int, err error) {
	return 80, 24, nil
}

func TestTerminalDimensions(t *testing.T) {
	// Replace the real Terminal variable with our mock implementation.
	Terminal = MockTerminal{}

	dimensions, err := TerminalDimensions()
	if err != nil {
		t.Errorf("TerminalDimensions() error = %v, wantErr nil", err)
		return
	}
	if dimensions == nil {
		t.Errorf("TerminalDimensions() = nil, want non-nil")
	}
	if len(dimensions) != 2 {
		t.Errorf("TerminalDimensions() length = %v, want 2", len(dimensions))
	}
}
