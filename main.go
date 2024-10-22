package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Printf("line %d, content: %v\n", i, scanner.Text())
		i++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
