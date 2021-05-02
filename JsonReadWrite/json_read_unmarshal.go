package main

import (
	"encoding/json"
	"io/ioutil"
)

func readJSONUnmarshal(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	datas := []map[string]interface{}{}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		// Do some filtering
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}
