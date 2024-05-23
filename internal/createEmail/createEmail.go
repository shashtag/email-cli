package createemail

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

func CreateEmail() (string, error) {
	file, err := os.Open("./combined.csv")
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer file.Close()

	inputIncidentNumbers()

	// Create a new CSV reader
	// reader := csv.NewReader(file)

	// Read the header row
	// header, err := reader.Read()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", nil
	// }

	//parse the html template file
	var html bytes.Buffer
	tmpl, err := template.ParseFiles(viper.GetString("ROOT_DIR") + "/html/wsgm.html")
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	tmpl.Execute(&html, nil)

	fmt.Println(html.String())

	// // Create a new HTML table
	// table := "<table><thead><tr>"
	// for _, column := range header {
	// 	table += fmt.Sprintf("<th>%s</th>", column)
	// }
	// table += "</tr></thead><tbody>"

	// // Read the data rows
	// for {
	// 	record, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return "", nil
	// 	}

	// 	// Add the row to the HTML table
	// 	table += ("<tr>")
	// 	for _, value := range record {
	// 		table += fmt.Sprintf("<td>%s</td>", value)
	// 	}
	// 	table += "</tr>"
	// }

	// // Close the HTML table
	// table += "</tbody></table>"

	// // Print the HTML table
	// fmt.Println(table)
	return html.String(), nil
}

func getNumberOfIncidents() int {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return len(files) - 1

}

func inputIncidentNumbers() (string, string) {

	incidentCount := getNumberOfIncidents()

	var firstIncident string
	var incidentList string

	for i := range incidentCount {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("Incident %d", i+1),
		}

		result, err := prompt.Run()
		if err == nil {
			pterm.Error.Println(err)
			fmt.Printf("pterm.Error.Fatal: %v\n", pterm.Error.Fatal)
		}

		if i == 0 {
			firstIncident = result
		}
		incidentList += result
		if i != incidentCount-1 {
			incidentList += " / "
		}

	}

	return firstIncident, incidentList

}
