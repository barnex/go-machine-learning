package vs

import (
	"testing"
)

// Test raw derivatives by comparing to numerical approximation
func TestDropout_Diff(t *testing.T) {
	d := Dropout(100, 2, 0.1)
	testDiffW(t, d)
	testDiffX(t, d)
	d.NextState()
	testDiffW(t, d)
	testDiffX(t, d)
	d.NextState()
	testDiffW(t, d)
	testDiffX(t, d)
}
