package createemail

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateEmail() error {
	file, err := os.Open("combined.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	fmt.Println(records)

	// Process the records and create the email

	return nil
}
