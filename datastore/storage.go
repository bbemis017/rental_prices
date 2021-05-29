package datastore

import (
	"strings"
)

type CSVStore struct {
	Data string // string in quoted csv format
}

func NewCSVStore(header []string) CSVStore {
	return CSVStore{
		Data: strings.Join(quoteStrings(header), ",") + "\n",
	}
}

// Quotes all values in string array in place
func quoteStrings(record []string) []string {
	for index := range record {
		record[index] = "\"" + record[index] + "\""
	}
	return record
}

// Writes a record to the CSV file
func (store *CSVStore) WriteRecord(record []string) {
	store.Data += strings.Join(quoteStrings(record), ",") + "\n"
}
