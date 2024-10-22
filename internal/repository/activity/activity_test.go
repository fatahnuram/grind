package activity_test

import (
	"fmt"
	"testing"

	"github.com/fatahnuram/grind/internal/repository/activity"
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
			got := activity.DayIntToString(suite.WantNum)
			if got != suite.WantString {
				t.Errorf("incorrect day int to string conversion, want: %v, got: %v", suite.WantString, got)
			}
		})
	}
}

func TestDayIntToStringMod7(t *testing.T) {
	suites := []struct {
		WantNum    int
		WantString string
	}{
		{7, "sun"},
		{15, "mon"},
		{23, "tue"},
		{31, "wed"},
		{39, "thu"},
		{47, "fri"},
		{55, "sat"},
	}

	for _, suite := range suites {
		name := fmt.Sprintf("%d-%v", suite.WantNum, suite.WantString)
		t.Run(name, func(t *testing.T) {
			got := activity.DayIntToString(suite.WantNum)
			if got != suite.WantString {
				t.Errorf("incorrect day int to string conversion, want: %v, got: %v", suite.WantString, got)
			}
		})
	}
}

func TestNewActivity(t *testing.T) {
	wantName := ""
	wantFreq := 0
	wantDay := 0
	wantDate := 0
	wantMonth := 0
	wantFunc := ""
	wantIsComplete := false
	wantCompletedAt := ""

	got := activity.NewActivity()

	if got.Name != wantName {
		t.Errorf("incorrect activity name, want: %v, got: %v", wantName, got.Name)
	}
	if got.Frequency != activity.Frequency(wantFreq) {
		t.Errorf("incorrect activity frequency, want: %d, got: %d", wantFreq, got.Frequency)
	}
	if got.Day != wantDay {
		t.Errorf("incorrect activity day, want: %d, got: %d", wantDay, got.Day)
	}
	if got.Date != wantDate {
		t.Errorf("incorrect activity date, want: %d, got: %d", wantDate, got.Date)
	}
	if got.Month != wantMonth {
		t.Errorf("incorrect activity month, want: %d, got: %d", wantMonth, got.Month)
	}
	if got.Function != wantFunc {
		t.Errorf("incorrect activity function, want: %v, got: %v", wantFunc, got.Function)
	}
	if got.IsCompleted != wantIsComplete {
		t.Errorf("incorrect activity complete status, want: %t, got: %t", wantIsComplete, got.IsCompleted)
	}
	if got.CompletedAt != wantCompletedAt {
		t.Errorf("incorrect activity complete timestamp, want: %v, got: %v", wantCompletedAt, got.CompletedAt)
	}
}
