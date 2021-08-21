package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
func main() {
	csvFilename := flag.String("csv", "problems.csv", "a .CSV file in the format of 'question,answer'")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided .CSV file.")
	}
	fmt.Println(lines)

}
