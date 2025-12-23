package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	f, err := os.Open("./student_scores.csv")
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
	// once we have all the data, we draw each graph
	for c, values := range columnsValues {
		// create a new plot
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of %s", records[0][c])
		// create a new normalized histogram
		// and add it to the plot
		h, err := plotter.NewHist(values, 16)
		if err != nil {
			log.Fatal(err)
		}
		h.Normalize(1)
		p.Add(h)
		// save the plot to a PNG file.
		if err := p.Save(
			10*vg.Centimeter,
			10*vg.Centimeter,
			"student_scores_hist.png",
		); err != nil {
			log.Fatal(err)
		}
	}
}
