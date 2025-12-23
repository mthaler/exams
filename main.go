package main

import (
	"encoding/csv"
	"log"
	"os"
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
	for i, record := range records {
	}
}
