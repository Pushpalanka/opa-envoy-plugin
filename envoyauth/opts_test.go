package envoyauth

import (
	"testing"
)

func TestNewEvalResultWithDecisionID(t *testing.T) {
	expectedDecisionID := "some-decision-id"

	er, _, err := NewEvalResult(WithDecisionID(expectedDecisionID))
	if err != nil {
		t.Fatalf("NewEvalResult() error = %v, wantErr %v", err, false)
	}
	if er.DecisionID != expectedDecisionID {
		t.Errorf("Expected DecisionID to be '%v', got '%v'", expectedDecisionID, er.DecisionID)
	}
}

func TestNewEvalResultWithoutOptDecisionID(t *testing.T) {
	_, _, err := NewEvalResult()
	if err != nil {
		t.Fatalf("NewEvalResult() error = %v, wantErr %v", err, false)
	}
}
