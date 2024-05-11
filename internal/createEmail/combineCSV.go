package createemail

import (
	"bufio"
	"fmt"
	"os"
)

func CombineCSV() error {
	// Get the list of files to combine.
	folder, err := os.Open("./") //open the current directory
	if err != nil {
		fmt.Println("error opening directory:", err) //print error if directory is not opened
		return err
	}
	defer folder.Close()                 //close the directory opened
	filesData, err := folder.Readdir(-1) //read the files from the directory
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return err
	}

	outputFile, err := os.Create("combined.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Combine the files.
	for i, filename := range filesData {
		// Open the file.
		file, err := os.Open(filename.Name())
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Create a new scanner to read the file line by line.
		scanner := bufio.NewScanner(file)

		// Skip the first line of the file, except for the first file.
		if i > 0 {
			scanner.Scan()
		}

		// Write the file data to the output file.
		for scanner.Scan() {
			_, err := outputFile.Write(scanner.Bytes())
			if err != nil {
				panic(err)
			}

			// Write a newline character to the output file.
			_, err = outputFile.Write([]byte("\n"))
			if err != nil {
				panic(err)
			}
		}
	}
	return nil
}
