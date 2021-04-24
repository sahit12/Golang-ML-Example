package main

import (
	rd "Golang-ML-Example/ReadHelper"
	"fmt"
)

func main() {
	filename := "Datasets/archive/sample.json"
	s := rd.LoadJSON(filename)
	fmt.Println(s.Versions[0].Version)
}
