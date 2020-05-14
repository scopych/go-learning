package main

// Reads csv file with servise report and counts how many servants in each group
import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"os"
	"regexp"
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

	group := 1
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			//      log.Fatal(err)
			panic(err)
		}
		// Check if record where 1-t colon is number, 2-d colon is first name and last name
		fio, _ := regexp.MatchString(`[А-Я](\p{Cyrillic})+ +[А-Я](\p{Cyrillic})+`, record[2])
		n, _ := regexp.MatchString(`[[:digit:]]+`, record[1])
		if n && fio {
			count := 0
			for count = 0; fio && n; count++ { // counts num of such records
				// Read each record from csv
				record, err = r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					//      log.Fatal(err)
					panic(err)
				}
				fio, _ = regexp.MatchString(`(\p{Cyrillic})+ +(\p{Cyrillic})+`, record[2])
				n, _ = regexp.MatchString(`[[:digit:]]+`, record[1])
			}
			fmt.Printf("Group %d: %d servants\n", group, count)
			group++
		}
	}
}
