package createemail

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func CreateEmail() error {
	file, err := os.Open("./combined.csv")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Create a new HTML table
	table := "<table><thead><tr>"
	for _, column := range header {
		table += fmt.Sprintf("<th>%s</th>", column)
	}
	table += "</tr></thead><tbody>"

	// Read the data rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Add the row to the HTML table
		table += ("<tr>")
		for _, value := range record {
			table += fmt.Sprintf("<td>%s</td>", value)
		}
		table += "</tr>"
	}

	// Close the HTML table
	table += "</tbody></table>"

	// Print the HTML table
	fmt.Println(table)
	return nil
}
