package vs

import (
	"testing"
)

// Test raw derivatives by comparing to numerical ≈imation
func TestActivation_DiffX(t *testing.T) {
	testDiffX(t, NewActivation(7, Re))
}
