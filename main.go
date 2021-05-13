package main

import (
	"Golang-ML-Example/ml"
	"Golang-ML-Example/utility"
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Initialize the utility settings

	s := &utility.Settings{
		Columns:  []string{"Weight", "Vertical_len", "Diagonal_len", "Cross_len", "Height", "Width"},
		DataFile: "datasets/Fish.csv",
		PlotType: "scatter",
	}

	// Load data from file
	DF := s.CleanAndLoad()

	// Load Plotting data
	for _, name := range s.Columns {
		pltv, pltxy := s.GetColumnData(DF, name, "Weight")

		p := plot.New()

		switch s.PlotType {
		case "hist":
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
			err = p.Save(200, 200, "images/"+name+"_histogram.png")
			if err != nil {
				log.Panic(err)
			}

		case "scatter":
			p.X.Label.Text = name
			p.Y.Label.Text = "Weight"

			p.Add(plotter.NewGrid())
			s, err := plotter.NewScatter(pltxy)
			if err != nil {
				log.Fatalf("Error after Scatter:: %v", err)
			}

			s.GlyphStyle.Radius = vg.Points(1)

			// Save the plot to a PNG file.
			p.Add(s)
			err = p.Save(4*vg.Inch, 4*vg.Inch, "images/"+name+"_scatter.png")
			if err != nil {
				log.Fatalf("Error while saving image:: %v", err)
			}
		}
	}

	// Create training and test data
	utility.TrainTestSplit(DF, true)

	// Train and get value
	rval := ml.TrainRegression("Selling Price", "seats", 0, "datasets/Fish.csv")
	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", rval.Formula)

}
