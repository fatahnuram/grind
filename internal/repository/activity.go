package repository

import (
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
	tempTime := time.Date(2009, 11, day+1, 0, 0, 0, 0, time.UTC)
	stringDay := tempTime.Format("Mon")
	return strings.ToLower(stringDay)
}

func DayStringToInt(day string) int {
	// TODO: implement
	return -1
}

func MonthIntToString(month int) string {
	// TODO: implement
	return "wip"
}

func MonthStringToInt(month string) int {
	// TODO: implement
	return -1
}
