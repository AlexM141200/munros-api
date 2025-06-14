package csv


import (
	"encoding/csv"
	"fmt"
	"os"
)

func openCsvFile(){
	file, err := os.Open("../../data/mumrotab_v8.0.1.csv")
	if err != nil {
		panic(err)
	}

	defer file.close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, row := range data{
		for _, col := range row{
			fmt.Printf("%s", col)

		}
		fmt.Println()
	}
}
