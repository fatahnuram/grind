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

	act := activity.NewActivity()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); {
		if i == 0 {
			// csv headers
			continue
		}
		// fmt.Printf("line %d, content: %v\n", i, scanner.Text())
		activity.CsvToActivity(scanner.Text(), &act)
		i++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
