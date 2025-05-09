package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"practice_01/model"
)

func LoadData(path string) model.Data {
	bytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	fmt.Println(bytes)

	var data model.Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("Error unmarshaling data: %v", err)
	}

	return data
}
