package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ReadCSV iterates over the records and print the content
func ReadCSV(r *csv.Reader) {
	// Ignore header
	_, _ = r.Read()
	// iterate over records
	for record, err := r.Read(); err != io.EOF; record, err = r.Read() {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s possui %s anos\n", record[0], record[1])
	}
}

// ReadAllCSV iterates over the records and print the content
func ReadAllCSV(r *csv.Reader) {
	// read all records
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// ignore header
	records = records[1:]
	// iterate over records
	for _, record := range records {
		fmt.Printf("%s possui %s anos\n", record[0], record[1])
	}
}

func main() {
	in := `name,age
Cássio,26
Pedro,24
Thiago,25
`
	// CSVReader based on a Reader
	// it's cool because can be anything that implements Reader interface
	r := csv.NewReader(strings.NewReader(in))
	println("\nComma == ,")
	ReadCSV(r)

	// Changing the input format
	in = `name;age
Cássio;26
Pedro;24
Thiago;25
`
	r = csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	println("\nComma == ;")
	ReadCSV(r)

	// Changing input format and add some comments
	in = `name|age
Cássio|26
# comments
Pedro|24
# comments
Thiago|25
`
	r = csv.NewReader(strings.NewReader(in))
	r.Comma = '|'
	// will ignore comments
	r.Comment = '#'
	println("\nComma == |")
	// Another method to read a csv reader
	ReadAllCSV(r)

	// records
	records := [][]string{
		{"name", "age"},
		{"Cássio", "26"},
		{"Pedro", "24"},
		{"Thiago", "24"},
	}

	// will write write csv on stdout
	w := csv.NewWriter(os.Stdout)
	println("\nWrite")

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	// simplest way
	println("\nWrite all")
	w2 := csv.NewWriter(os.Stdout)
	w2.WriteAll(records)
	w2.Flush()

}
