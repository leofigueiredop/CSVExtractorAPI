package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

var People []models.Person

func LoadCSVs() {
	csvfile, err := os.Open("")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.Comma = ';'
	r.LazyQuotes = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		p := models.Person{
			Name: record[0],
			CPF:  record[1],
			City: record[2],
		}
		People = append(People, p)
	}
}

func Search(term string, layout string) []interface{} {
	switch layout {
	case "pessoa":
		return searchInPeople(term)
	default:
		return nil
	}
}

func searchInPeople(term string) []interface{} {
	result := make([]interface{}, 0)

	for _, p := range People {
		if strings.Contains(p.Name, term) {
			result = append(result, p)
		}
	}
	return result
}
