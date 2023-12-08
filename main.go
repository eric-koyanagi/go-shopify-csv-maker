package main

import (
	"csvMaker/getDummyRows"
	"encoding/csv"
	"fmt"
	"os"
)

const ROWS_TO_GENERATE = 100000
const OUTPUT_FILE = "example_products.csv"

func main() {

	//writer, err := openNewCSV()
	file, err := os.Create(OUTPUT_FILE)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer file.Close()

	// create the writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the headers
	err = writeRow(writer, getDummyRows.Headers)
	if err != nil {
		fmt.Println("Error writing header row", err)
	}

	// create the dummy data
	for i := 0; i < ROWS_TO_GENERATE; i++ {
		row := getDummyRows.NewRowAsSlice("")

		err = writeRow(writer, row)
		if err != nil {
			fmt.Println("Error writing data row", err)
		}

		// creates variants, which use the same handle
		if getDummyRows.RandomlyHasVariants() {
			numVariants := getDummyRows.RandomVariantNum()
			for x := 0; x+i < ROWS_TO_GENERATE && x < numVariants; x++ {
				// because row is ordered, we know row[0] is always the handle
				variantRow := getDummyRows.NewRowAsSlice(row[0])
				err = writeRow(writer, variantRow)
				if err != nil {
					fmt.Println("Error writing variant row", err)
				}
			}

			i += numVariants
		}
	}

	fmt.Printf("Done! Created %d products", ROWS_TO_GENERATE)
}

func writeRow(writer *csv.Writer, row []string) error {
	err := writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}
