package services

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Record struct {
	Type     string `json:"type"`
	CPF      string `json:"CPFV"`
	FullName string `json:"full_name"`
	Status   string `json:"status"`
	Link     string `json:"link"`
}

func Convert(imputDir string, outputFileName string) {
	directory := imputDir
	var allRecords []Record
	recordMap := make(map[string]struct{}) // to keep track of duplicates

	files, err := os.ReadDir(imputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".csv" {
			filePath := filepath.Join(imputDir, file.Name())
			file, err := os.Open(filePath)
			if err != nil {
				log.Println("Unable to read input file: %s", err)
			}

			defer file.Close()

			reader := csv.NewReader(file)
			reader.FieldsPerRecord = -1
			reader.Read() // skip header line

			for {
				line, err := reader.Read()
				if err == io.EOF {
					log.Println("Unable to read input file: %s", err)
					break
				}
				record := Record{}
				if len(line) == 6 {
					record = Record{
						Type:     line[0],
						CPF:      line[2],
						FullName: line[3],
						Status:   line[4],
						Link:     line[5],
					}
				} else if len(line) == 7 {
					record = Record{
						Type:     line[0],
						CPF:      line[2],
						FullName: line[3],
						Status:   line[4] + " - " + line[5],
						Link:     line[6],
					}
				} else if len(line) > 7 {
					fmt.Printf("Registro maior que 7: %v\n", record)
				}

				// Check for duplicates before appending
				if _, exist := recordMap[record.CPF]; exist {
					fmt.Printf("Duplicate record found: %v\n", record)
				} else {
					recordMap[record.CPF] = struct{}{}
					allRecords = append(allRecords, record)
				}
			}
		}
	}

	jsonData, err := json.Marshal(allRecords)
	if err != nil {
		log.Println("Unable to read input file: %s", err)
	}

	jsonFile, err := os.Create(directory + outputFileName)
	if err != nil {
		log.Println("Unable to read input file: %s", err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

}
