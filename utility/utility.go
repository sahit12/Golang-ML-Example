package utility

import (
	"log"
	"math"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot/plotter"
)

type Settings struct {
	Columns  []string
	DataFile string
	PlotType string
}

func (s *Settings) CleanAndLoad() dataframe.DataFrame {
	log.Println("Cleaning Started")

	// Open the CSV file.
	dataFile, err := os.Open(s.DataFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dataFile.Close()

	// Create a dataframe from the CSV file.
	DF := dataframe.ReadCSV(dataFile)
	log.Println("Data Loaded")

	return DF
}

func (s *Settings) GetColumnData(
	df dataframe.DataFrame,
	colname string,
	XYname string) (*plotter.Values, *plotter.XYs) {

	// Create a plotter.Values value and fill it with the
	// values from the required column of the dataframe.
	switch s.PlotType {
	case "hist":
		plotvalues := make(plotter.Values, df.Nrow())

		for i, value := range df.Col(colname).Float() {
			if !math.IsNaN(value) {
				plotvalues[i] = value
			} else {
				log.Printf("Passing on with the value at %v", i)
				continue
			}
		}
		return &plotvalues, nil
	case "scatter":
		y := df.Col(XYname).Float()
		points := make(plotter.XYs, df.Nrow())

		for i, value := range df.Col(colname).Float() {
			if !math.IsNaN(value) && !math.IsNaN(y[i]) {
				points[i].X = value
				points[i].Y = y[i]
			} else {
				log.Printf("Handling data at the value %v", i)
				continue
			}
		}
		return nil, &points
	}
	return nil, nil
}
