package ReadHelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ArxivMetaData struct {
	ID            string     `json:"id"`
	Submitter     string     `json:"submitter"`
	Authors       string     `json:"authors"`
	Title         string     `json:"title"`
	Comments      string     `json:"comments"`
	JournalRef    string     `json:"journal-ref"`
	Doi           string     `json:"doi"`
	ReportNo      string     `json:"report-no"`
	Categories    string     `json:"categories"`
	License       string     `json:"license"`
	Abstract      string     `json:"abstract"`
	Versions      []Versions `json:"versions"`
	UpdateDate    string     `json:"update_date"`
	AuthorsParsed [][]string `json:"authors_parsed"`
}

type Versions struct {
	Version string `json:"version"`
	Created string `json:"created"`
}

func LoadJSON(filename string) ArxivMetaData {
	// Open the data file to start reading
	datafile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: [%v]", err)
	}
	defer datafile.Close()

	body, err := ioutil.ReadAll(datafile)
	if err != nil {
		log.Fatalf("Error: [%v]", err)
	}

	// Declare a variable of type ArxivMetaData
	var amt ArxivMetaData

	// Unmarshal JSON data onto the variable
	if err := json.Unmarshal(body, &amt); err != nil {
		log.Fatalf("Error: [%v]", err)
	}

	return amt
}
