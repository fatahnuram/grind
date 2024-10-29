package activity

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
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
var Months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
var Frequencies = []string{"daily", "weekly", "monthly", "annually", "custom"}

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
	if day < 0 {
		return ""
	}
	tempTime := time.Date(2009, 11, (day%7)+1, 0, 0, 0, 0, time.UTC)
	stringDay := tempTime.Format("Mon")
	return strings.ToLower(stringDay)
}

func DayStringToInt(day string) int {
	return slices.Index(Days, day)
}

func MonthStringToInt(month string) int {
	result := slices.Index(Months, month)
	if result < 0 {
		return result
	} else {
		return result + 1
	}
}

func MonthIntToString(month int) string {
	if month < 1 || month > 12 {
		return ""
	}
	return Months[month-1]
}

func FreqStringToInt(freq string) Frequency {
	return Frequency(slices.Index(Frequencies, freq))
}

func FreqIntToString(freq Frequency) string {
	if int(freq) >= len(Frequencies) || int(freq) < 0 {
		return ""
	}
	return Frequencies[freq]
}

func DateIntToString(d int) string {
	if d < 1 {
		return ""
	}
	return fmt.Sprintf("%d", d)
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

	switch act.Frequency {
	case Daily:
		act.Day = -1
		act.Date = -1
		act.Month = -1
		act.Function = ""

	case Weekly:
		act.Day = DayStringToInt(cols[2])
		act.Date = -1
		act.Month = -1
		act.Function = ""

	case Monthly:
		act.Day = -1
		act.Date = strToInt(cols[3])
		act.Month = -1
		act.Function = ""

	case Annually:
		act.Day = -1
		act.Date = strToInt(cols[3])
		act.Month = MonthStringToInt(cols[4])
		act.Function = ""

	case Custom:
		act.Day = -1
		act.Date = -1
		act.Month = -1
		act.Function = cols[5]

	default:
		log.Fatalf("cannot parse frequency: %v", cols[1])
	}
}

func PrettyPrint(acts []Activity) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	// headers
	fmt.Fprintln(w, strings.ToUpper("name\tfrequency\tday\tdate\tmonth\tfunction"))

	for _, act := range acts {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", act.Name, FreqIntToString(act.Frequency), DayIntToString(act.Day), DateIntToString(act.Date), MonthIntToString(act.Month), act.Function)
	}
	w.Flush()
}
