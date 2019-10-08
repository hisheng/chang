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
		fmt.Println("INSERT IGNORE INTO wkread.smartprogram_toufang_feeds (feed_id,start_day,end_day,status,toufang_channel) VALUES ("+row[0]+",'2019-09-17','2019-12-17',0,0);")
	}


}
