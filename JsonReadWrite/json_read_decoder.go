package main

import (
	"encoding/json"
	"os"
)

func readJSONDecoder(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)

	filteredData := []map[string]interface{}{}

	// Read the array open bracket
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)

		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}
