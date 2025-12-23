package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
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
	//exams.FieldsPerRecord = 2
	records, err := exams.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// by slicing the records we skip the header
	records = records[1:]
	columnsValues := map[int]plotter.Values{}
	for i, record := range records {
		for c := 0; c < exams.FieldsPerRecord; c++ {
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

// resultXY --> sum((x-meanX)*(y-meanY))
// resultXX --> sum((x-meanX)^2)
func sumXYandXX(arrayX []float64, arrayY []float64, meanX float64, meanY float64) (float64, float64) {
	resultXX := 0.0
	resultXY := 0.0
	for x := 0; x < len(arrayX); x++ {
		for y := 0; y < len(arrayY); y++ {
			if x == y {
				resultXY += (arrayX[x] - meanX) * (arrayY[y] - meanY)
			}
		}
		resultXX += (arrayX[x] - meanX) * (arrayX[x] - meanX)
	}
	return resultXY, resultXX
}

// estimateBoB1 --> Function that calculates the regression coefficients b0 and b1
// y_predicted = b0 + b1*x_input
func estimateB0B1(x []float64, y []float64) (float64, float64) {
	var meanX float64
	var meanY float64
	var sumXY float64
	var sumXX float64

	meanX = stat.Mean(x, nil) //mean of x
	meanY = stat.Mean(y, nil) //mean pf ysumXY, sumXX = sumXYandXX(x, y, meanX, meanY)// regression coefficients
	b1 := sumXY / sumXX       // b1 or slope
	b0 := meanY - b1*meanX    // b0 or interceptreturn b0, b1
}
