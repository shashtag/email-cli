package createemail

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"

	"github.com/pterm/pterm"
)

func CombineCSV() error {
	// Get the list of files to combine.
	filesData := getFilesInFolder()

	outputFile, err := os.Create("combined.csv")
	if err != nil {
		return fmt.Errorf("error creating output file : %w", err)
	}
	defer outputFile.Close()

	// Combine the files.
	for i, filename := range filesData {
		// Skip the output file.
		if filename.Name() == "combined.csv" {
			continue
		}

		file := openFile(filename.Name())
		defer file.Close()

		// Create a new scanner to read the file line by line.
		scanner := bufio.NewScanner(file)

		// Skip the first line of the file
		if i > 0 {
			scanner.Scan()
		}

		// Write the file data to the output file.
		for scanner.Scan() {
			_, err := outputFile.Write(scanner.Bytes())
			if err != nil {
				pterm.Fatal.Println("error writing to output file : ", err)
			}

			// Write a newline character to the output file.
			_, err = outputFile.Write([]byte("\n"))
			if err != nil {
				pterm.Fatal.Println("error writing to output file : ", err)
			}
		}
	}
	return nil
}

func getFilesInFolder() []fs.FileInfo {
	folder, err := os.Open("./") //open the current directory
	if err != nil {
		pterm.Fatal.Println("error opening directory:", err)
	}

	defer folder.Close() //close the directory opened

	filesData, err := folder.Readdir(-1) //read the files from the directory
	if err != nil {
		pterm.Fatal.Println("error reading directory:", err)

	}

	return filesData
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		pterm.Fatal.Println("error opening file : ", filename, err)
	}
	return file
}
