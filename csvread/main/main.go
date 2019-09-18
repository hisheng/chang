package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main()  {
	csvFile, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		// process the `row` here
		fmt.Println(row)
	}


}
