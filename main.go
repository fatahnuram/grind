package main

import (
	"bufio"
	"fmt"
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
	for i := 0; scanner.Scan(); {
		if i == 0 {
			// csv headers
			i++
			continue
		}
		// fmt.Printf("line %d, content: %v\n", i, scanner.Text())
		act := activity.NewActivity()
		activity.CsvToActivity(scanner.Text(), &act)
		fmt.Printf("activity: %#v\n", act)
		i++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
