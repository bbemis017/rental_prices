package datastore

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVStore struct {
	Filepath string
	Length   int
}

func NewCSVStore(filepath string) CSVStore {
	length := 0

	file, openError := os.Open(filepath)
	if openError == nil {
		filedata, err := csv.NewReader(file).ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		length = len(filedata)
	}
	defer file.Close()

	return CSVStore{
		Filepath: filepath,
		Length:   length,
	}
}

// Writes a record to the CSV file
func (store *CSVStore) WriteRecords(records [][]string) error {
	file, openError := os.OpenFile(store.Filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if openError != nil {
		return openError
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(records)

	store.Length += len(records)

	return nil
}

// Writes a record to the CSV file
func (store *CSVStore) WriteRecord(record []string) error {
	file, openError := os.OpenFile(store.Filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if openError != nil {
		return openError
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(record)

	store.Length++

	return nil
}
