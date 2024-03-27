package services

import (
	"encoding/json"
	"log"
	"os"
)

type PEPS struct {
	CODCONTROLE string `json:"CODCONTROLE"`
	DATASET     string `json:"DATASET"`
	TIPO        string `json:"TIPO"`
	CPFCNPJ     string `json:"CPFCNPJ"`
	NOME        string `json:"NOME"`
	Data        string `json:"Data"`
	AUX         string
}

type CPFS struct {
	Type     string `json:"type"`
	CPFV     string `json:"CPFV"`
	FullName string `json:"full_name"`
	Status   string `json:"status"`
	Link     string `json:"link"`
}

func MatchPepsNames() {
	var data1 []PEPS
	var data2 []CPFS

	// Read JSON 1
	jsonRaw, err := os.ReadFile("resultPEPPF.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonRaw, &data1)

	// Read JSON 2 and create a map for easier searching
	jsonRaw, err = os.ReadFile("files/PF/cpfsallCpfsvm17h42m.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonRaw, &data2)
	data2Map := make(map[string]CPFS)
	for _, record := range data2 {
		data2Map[record.CPFV] = record
	}

	var foundSameName []PEPS
	var foundDiffName []PEPS
	var notFound []PEPS

	for _, record1 := range data1 {
		record2, found := data2Map[record1.CPFCNPJ]
		if found && record1.NOME == record2.FullName {
			foundSameName = append(foundSameName, record1)
		} else if found {
			record1.AUX = record2.FullName
			foundDiffName = append(foundDiffName, record1)
		} else {
			notFound = append(notFound, record1)
		}
	}

	// Write result files
	jsonRaw, err = json.MarshalIndent(foundSameName, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("found_same_name.json", jsonRaw, 0644)
	if err != nil {
		log.Fatal(err)
	}

	jsonRaw, err = json.MarshalIndent(foundDiffName, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("found_diff_name.json", jsonRaw, 0644)
	if err != nil {
		log.Fatal(err)
	}

	jsonRaw, err = json.MarshalIndent(notFound, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("not_found.json", jsonRaw, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
