package ReadHelper

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadCSV(filename string) [][]string {

	// Open the data file to start reading
	datafile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: [%v]", err)
	}
	defer datafile.Close()

	// We create a new csv reader to start
	// reading from the opened file
	reader := csv.NewReader(datafile)

	// In case we are not aware of the number
	// of fields per row, we can use below function
	reader.FieldsPerRecord = -1

	// Read all csv records
	rawcsvdata, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error: [%v]", err)
	}

	return rawcsvdata
}
