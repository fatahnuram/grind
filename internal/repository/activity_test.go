package repository_test

import (
	"fmt"
	"testing"

	"github.com/fatahnuram/grind/internal/repository"
)

func TestDayIntToString(t *testing.T) {
	suites := []struct {
		WantNum    int
		WantString string
	}{
		{0, "sun"},
		{1, "mon"},
		{2, "tue"},
		{3, "wed"},
		{4, "thu"},
		{5, "fri"},
		{6, "sat"},
	}

	for _, suite := range suites {
		name := fmt.Sprintf("%d-%v", suite.WantNum, suite.WantString)
		t.Run(name, func(t *testing.T) {
			got := repository.DayIntToString(suite.WantNum)
			if got != suite.WantString {
				t.Errorf("incorrect day int to string conversion, want: %v, got: %v", suite.WantString, got)
			}
		})
	}
}
