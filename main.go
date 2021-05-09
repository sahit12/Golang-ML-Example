package main

import (
	"Golang-ML-Example/utility"
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {

	// Initialize the utility settings

	s := &utility.Settings{
		Columns:  []string{"selling_price", "km_driven", "seats"},
		DataFile: "Datasets/car_data/Car details v3.csv",
	}

	// Load data from file
	DF := s.CleanAndLoad()
	// fmt.Println(DF.Describe())

	// Load Plotting data
	for _, name := range s.Columns {
		pltv := s.GetColumnData(DF, name)

		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of %s", name)

		// Create a histogram of our values drawn
		// from the standard normal.
		h, err := plotter.NewHist(pltv, 5)
		if err != nil {
			log.Println(err)
		}

		// Normalize the histogram.
		h.Normalize(1)

		// Add the histogram to the plot.
		p.Add(h)
		err = p.Save(200, 200, "Images/"+name+"_histogram.png")
		if err != nil {
			log.Panic(err)
		}
	}
}
