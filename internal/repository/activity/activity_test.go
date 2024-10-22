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

func TestFreqStringToInt(t *testing.T) {
	suites := []struct {
		Freq string
		Want activity.Frequency
	}{
		{Freq: "daily", Want: activity.Daily},
		{Freq: "weekly", Want: activity.Weekly},
		{Freq: "monthly", Want: activity.Monthly},
		{Freq: "annually", Want: activity.Annually},
		{Freq: "custom", Want: activity.Custom},
		{Freq: "xyz", Want: -1},
	}

	for _, suite := range suites {
		t.Run(suite.Freq, func(t *testing.T) {
			got := activity.FreqStringToInt(suite.Freq)
			if got != suite.Want {
				t.Errorf("incorrect frequency string conversion, want: %v, got: %v", suite.Want, got)
			}
		})
	}
}

func TestDayStringToInt(t *testing.T) {
	suites := []struct {
		Name string
		Want int
	}{
		{Name: "sun", Want: 0},
		{Name: "mon", Want: 1},
		{Name: "tue", Want: 2},
		{Name: "wed", Want: 3},
		{Name: "thu", Want: 4},
		{Name: "fri", Want: 5},
		{Name: "sat", Want: 6},
		{Name: "xyz", Want: -1},
	}

	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			got := activity.DayStringToInt(suite.Name)

			if got != suite.Want {
				t.Errorf("incorrect day string to int conversion for %v, want: %d, got: %d", suite.Name, suite.Want, got)
			}
		})
	}
}

func TestMonthStringToInt(t *testing.T) {
	suites := []struct {
		Name string
		Want int
	}{
		{Name: "jan", Want: 1},
		{Name: "feb", Want: 2},
		{Name: "mar", Want: 3},
		{Name: "apr", Want: 4},
		{Name: "may", Want: 5},
		{Name: "jun", Want: 6},
		{Name: "jul", Want: 7},
		{Name: "aug", Want: 8},
		{Name: "sep", Want: 9},
		{Name: "oct", Want: 10},
		{Name: "nov", Want: 11},
		{Name: "dec", Want: 12},
		{Name: "xyz", Want: -1},
	}

	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			got := activity.MonthStringToInt(suite.Name)

			if got != suite.Want {
				t.Errorf("incorrect month string to int conversion for %v, want: %d, got: %d", suite.Name, suite.Want, got)
			}
		})
	}
}
