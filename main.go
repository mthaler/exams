package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot/plotter"
)

func main() {
	f, err := os.Open("./datasets/kc_house_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	exams := csv.NewReader(f)
	records, err := exams.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// by slicing the records we skip the header
	records = records[1:]
	columnsValues := map[int]plotter.Values{}
	for i, record := range records {
		for c := 2; c < exams.FieldsPerRecord; c++ {
			if _, found := columnsValues[c]; !found {
				columnsValues[c] = make(plotter.Values, len(records))
			}
			// we parse each close value and add it to our set
			floatVal, err := strconv.ParseFloat(record[c], 64)
			if err != nil {
				log.Fatal(err)
			}
			columnsValues[c][i] = floatVal
		}
	}
}
