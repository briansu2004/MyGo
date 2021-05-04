package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func makeData(n int) interface{} {
	datas := []map[string]interface{}{}
	for i := 0; i < n; i++ {
		data := map[string]interface{}{
			"Name":        fmt.Sprintf("Person %v", i),
			"Address":     "Somewhere we don't really know",
			"Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			"Age":         i + 10,
			"PhoneNumber": i + 99,
			"CreatedOn":   time.Now(),
		}
		datas = append(datas, data)
	}

	return datas
}

func BenchmarkMarshalWriteBig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		jsonString, _ := json.Marshal(makeData(1000000))
		ioutil.WriteFile("marhsall_big.json", jsonString, os.ModePerm)
	}
}

func BenchmarkEncodeWriteBig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file, _ := os.OpenFile("encode_big.json", os.O_CREATE, os.ModePerm)
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(makeData(1000000))
	}
}

func BenchmarkMarshalWriteSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		jsonString, _ := json.Marshal(makeData(10000))
		ioutil.WriteFile("marhsall_small.json", jsonString, os.ModePerm)
	}
}

func BenchmarkEncodeWriteSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file, _ := os.OpenFile("encode_small.json", os.O_CREATE, os.ModePerm)
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(makeData(10000))
	}
}
