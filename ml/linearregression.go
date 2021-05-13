package ml

import (
	"Golang-ML-Example/ReadHelper"
	"log"
	"strconv"

	"github.com/sajari/regression"
)

func TrainRegression(y string, m string, c int, trainset string) *regression.Regression {
	var r regression.Regression
	r.SetObserved(y)
	r.SetVar(c, m)

	traindata := ReadHelper.LoadCSV(trainset)

	for i, record := range traindata {

		// Skip the header
		if i == 0 {
			continue
		}

		// Parse the Weight column
		yval, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatalf("Error while parsing y value: %v", err)
		}
		// Parse the  column
		seatsVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yval, []float64{seatsVal}))
	}

	r.Run()

	return &r
}
