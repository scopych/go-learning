package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			//	log.Fatal(err)
			panic(err)
		}
		fmt.Printf("%2s %-20s %5s %9s %9s %9s %9s %-20s %-16s\n", record[1], record[2], record[3],
			record[4], record[5], record[6], record[7], record[8], record[9])
	}
}
