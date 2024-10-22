package helpers_test

import (
	"testing"

	"github.com/fatahnuram/grind/internal/helpers"
)

func TestIsBetween(t *testing.T) {
	suites := []struct {
		Name   string
		Target int
		Min    int
		Max    int
		Result bool
	}{
		{Name: "less than min", Target: 5, Min: 10, Max: 20, Result: false},
		{Name: "equal to min", Target: 10, Min: 10, Max: 20, Result: true},
		{Name: "middle", Target: 15, Min: 10, Max: 20, Result: true},
		{Name: "equal to max", Target: 20, Min: 10, Max: 20, Result: true},
		{Name: "more than max", Target: 25, Min: 10, Max: 20, Result: false},
	}

	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			got := helpers.IsBetween(suite.Target, suite.Min, suite.Max)
			if got != suite.Result {
				t.Errorf("incorrect evaluation %d between %d-%d, want: %t, got: %t", suite.Target, suite.Min, suite.Max, suite.Result, got)
			}
		})
	}
}
