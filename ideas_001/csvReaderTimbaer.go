package main

import (
	"flag"
	"fmt"
	"io"
	"encoding/csv"
	"io/ioutil"
	"strings"
)

const banner = "---------------------------------------------------------------------------------------------------------------------"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Row struct {
	columns []string
}

func (r Row) PrintRow() {
	fmt.Print("\n| ")
	for _, col := range r.columns {
		fmt.Printf(" %-20v |", col)
	}
}

type Csv struct {
	rows []Row
}

func (c Csv) Print() {
	fmt.Print("\n\n",banner)
	for _, row := range c.rows {
		row.PrintRow();
	}
	fmt.Print("\n", banner, "\n")
}

func (c Csv) Read(path string) Csv {
	fmt.Printf("Reading file from path: \"%v\".", path);

	in, err := ioutil.ReadFile(path);

	check(err);

	r := csv.NewReader(strings.NewReader(string(in)))
	r.TrimLeadingSpace = true
	r.Comma = ';'

	var csv Csv

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err);

		csv.rows = append(csv.rows, Row{record})
	}

	return csv
}

func main() {
	var (
		path string
		csv Csv
	)
	flag.StringVar(&path, "path", "ideas_001/books.csv", "Path to csv file")
	flag.Parse()

	csv = csv.Read(path)

	csv.Print()
}
