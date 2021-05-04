package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

const (
	fileIn  = "in.csv"
	fileOut = "out.csv"
)

type Employee struct {
	EmpNo     string
	FirstName string
	LastName  string
	Age       int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Display() {
	outFile, err := os.Create(fileOut)
	//outFile, err = os.OpenFile(fileOut, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)
	defer outFile.Close()

	inFile, err := os.Open(fileIn)
	if err != nil {
		fmt.Println(err.Error() + ": " + fileIn)
		return
	}
	defer inFile.Close()

	tmpl, err := template.New("WriteToFile").Parse("Employee {{.EmpNo}}\nName: {{.FirstName}} {{.LastName}}\nAge: {{.Age}}-year-old\n")
	check(err)

	r := csv.NewReader(inFile)
	check(err)

	r.Comma = ','
	r.Comment = '#'
	lines, err := r.ReadAll()
	check(err)

	//employees := make(map[string]*Employee, len(lines))
	n := 0
	for _, line := range lines {
		if n == 0 {
			//skip the header
			n++
			continue
		}

		age, err := strconv.Atoi(line[3])
		if err != nil {
			age = 0
		}
		employee := Employee{line[0], line[1], line[2], age}
		//employees[line[0]] = &employee

		err = tmpl.Execute(os.Stdout, employee)
		check(err)
	}
}

func WriteToFile() {
	outFile, err := os.Create(fileOut)
	//outFile, err = os.OpenFile(fileOut, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)
	defer outFile.Close()

	inFile, err := os.Open(fileIn)
	if err != nil {
		fmt.Println(err.Error() + ": " + fileIn)
		return
	}
	defer inFile.Close()

	tmpl, err := template.New("WriteToFile").Parse("Employee {{.EmpNo}}\nName: {{.FirstName}} {{.LastName}}\nAge: {{.Age}}-year-old\n")
	check(err)

	r := csv.NewReader(inFile)
	check(err)

	r.Comma = ','
	r.Comment = '#'
	lines, err := r.ReadAll()
	check(err)

	n := 0
	for _, line := range lines {
		if n == 0 {
			//skip the header
			n++
			continue
		}

		age, err := strconv.Atoi(line[3])
		if err != nil {
			age = 0
		}
		employee := Employee{line[0], line[1], line[2], age}

		err = tmpl.Execute(outFile, employee)
		check(err)
	}
}

func WriteToJson() {
	inFile, err := os.Open(fileIn)
	if err != nil {
		fmt.Println(err.Error() + ": " + fileIn)
		return
	}
	defer inFile.Close()

	r := csv.NewReader(inFile)
	check(err)

	r.Comma = ','
	r.Comment = '#'
	lines, err := r.ReadAll()
	check(err)

	file, _ := os.OpenFile("encode_out.json", os.O_CREATE, os.ModePerm)
	defer file.Close()

	employees := make(map[string]*Employee, len(lines))
	n := 0
	for _, line := range lines {
		if n == 0 {
			//skip the header
			n++
			continue
		}

		age, err := strconv.Atoi(line[3])
		if err != nil {
			age = 0
		}
		employee := Employee{line[0], line[1], line[2], age}
		employees[line[0]] = &employee

		check(err)
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(employees)

}

func main() {
	Display()

	WriteToFile()

	WriteToJson()
}
