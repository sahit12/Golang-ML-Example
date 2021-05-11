package main

import (
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
		Columns:  []string{"selling_price", "km_driven", "seats"},
		DataFile: "Datasets/car_data/Car details v3.csv",
		PlotType: "scatter",
	}

	// Load data from file
	DF := s.CleanAndLoad()
	// fmt.Println(DF.Describe())

	// Load Plotting data
	for _, name := range s.Columns {
		pltv, pltxy := s.GetColumnData(DF, name, "selling_price")
		// fmt.Println(pltxy)

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
			err = p.Save(200, 200, "Images/"+name+"_histogram.png")
			if err != nil {
				log.Panic(err)
			}

		case "scatter":
			p.X.Label.Text = "Selling Price"
			p.Y.Label.Text = name

			p.Add(plotter.NewGrid())
			s, err := plotter.NewScatter(pltxy)
			if err != nil {
				log.Fatalf("Error after Scatter:: %v", err)
			}

			s.GlyphStyle.Radius = vg.Points(1)

			// Save the plot to a PNG file.
			p.Add(s)
			err = p.Save(4*vg.Inch, 4*vg.Inch, "Images/"+name+"_scatter.png")
			if err != nil {
				log.Fatalf("Error while saving image:: %v", err)
			}
		}
	}
}
