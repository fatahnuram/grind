package activity

import (
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Frequency int

const (
	Daily Frequency = iota
	Weekly
	Monthly
	Annually
	Custom
)

const ActivityColumnCount = 6

var Days = []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}

type Activity struct {
	// from data csv
	Name      string
	Frequency Frequency
	Day       int
	Date      int
	Month     int
	Function  string

	// for internal handling
	IsCompleted bool
	CompletedAt string
}

func DayIntToString(day int) string {
	tempTime := time.Date(2009, 11, (day%7)+1, 0, 0, 0, 0, time.UTC)
	stringDay := tempTime.Format("Mon")
	return strings.ToLower(stringDay)
}

func DayStringToInt(day string) int {
	return slices.Index(Days, day)
}

func FreqStringToInt(freq string) Frequency {
	switch freq {
	case "daily":
		return Daily
	case "weekly":
		return Weekly
	case "monthly":
		return Monthly
	case "annually":
		return Annually
	case "custom":
		return Custom
	default:
		return -1
	}
}

func NewActivity() Activity {
	return Activity{
		IsCompleted: false,
		CompletedAt: "",
	}
}

func trim(txt string) string {
	return strings.TrimSpace(txt)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert to number: %v", s)
	}
	return i
}

func CsvToActivity(line string, act *Activity) {
	cols := strings.Split(line, ",")

	if len(cols) != ActivityColumnCount {
		log.Fatalf("column number does not match for \"%v\"", line)
	}

	for i := range cols {
		cols[i] = trim(cols[i])
	}

	act.Name = cols[0]
	act.Frequency = FreqStringToInt(cols[1])
	act.Day = DayStringToInt(cols[2])
	act.Date = strToInt(cols[3])
	// act.Month
	act.Function = cols[5]

	// walk at least 30 mins, daily, 0, 0, 0, 0
	// swimming 1 hour, weekly, sat, 0, 0, 0
	// pay internet and water bills, monthly, 0, 12, 0, 0
	// report annual taxes, annually, 0, 1, mar, 0
	// clean up cat's litter box, custom, 0, 0, 0, date%3==0
}
