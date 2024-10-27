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
	act.Month = MonthStringToInt(cols[4])
	act.Function = cols[5]

	// walk at least 30 mins, daily, 0, 0, 0, 0
	// swimming 1 hour, weekly, sat, 0, 0, 0
	// pay internet and water bills, monthly, 0, 12, 0, 0
	// report annual taxes, annually, 0, 1, mar, 0
	// clean up cat's litter box, custom, 0, 0, 0, date%3==0
}

func PrettyPrint(acts []Activity) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	// headers
	fmt.Fprintln(w, strings.ToUpper("name\tfrequency\tday\tdate\tmonth\tfunction"))

	for _, act := range acts {
		fmt.Fprintf(w, "%v\t%v\t%v\t%d\t%v\t%v\n", act.Name, FreqIntToString(act.Frequency), DayIntToString(act.Day), act.Date, MonthIntToString(act.Month), act.Function)
	}
	w.Flush()
}
