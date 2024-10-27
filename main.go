package main

import (
	"bufio"
	"log"
	"os"

	"github.com/fatahnuram/grind/internal/repository/activity"
)

var datafile = "./grind.csv"

func main() {
	file, err := os.Open(datafile)
	if err != nil {
		log.Fatalf("error opening file: %v, msg: %v", datafile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	activities := []activity.Activity{}
	for i := 0; scanner.Scan(); {
		if i == 0 {
			// csv headers
			i++
			continue
		}

		act := activity.NewActivity()
		activity.CsvToActivity(scanner.Text(), &act)
		activities = append(activities, act)
		i++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	activity.PrettyPrint(activities)
}
