package utility

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/kniren/gota/dataframe"
)

func TrainTestSplit(df dataframe.DataFrame, randomly bool) {
	// Calculate the number of elements in each set.
	trainingNum := (4 * df.Nrow()) / 5
	testNum := df.Nrow() - trainingNum

	// Create the subset indices.
	trainingIds := make([]int, trainingNum)
	testIds := make([]int, testNum)

	// Enumerate the training indices.
	for i := 0; i < trainingNum; i++ {
		trainingIds[i] = i
	}

	// Enumerate the test indices.
	for i := 0; i < testNum; i++ {
		testIds[i] = trainingNum + i
	}

	// Randomize both indices
	if randomly {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(trainingIds), func(i, j int) {
			trainingIds[i], trainingIds[j] = trainingIds[j], trainingIds[i]
		})
	}
	fmt.Println(trainingIds)

	// Create the subset dataframes.
	trainDF := df.Subset(trainingIds)
	testDF := df.Subset(testIds)

	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainDF,
		1: testDF,
	}

	// Create the respective files.
	for idx, setName := range []string{"train.csv", "test.csv"} {

		// Save the filtered dataset file.
		f, err := os.Create(path.Join("Datasets/", setName))
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer.
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
