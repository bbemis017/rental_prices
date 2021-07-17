package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that mapJsonToCsv returns results in header order
func TestMapJsonToCsvInHeaderOrder(t *testing.T) {

	headers := []string{"two", "one"}

	jsonData := map[string]interface{}{
		"one": "val1",
		"two": "val2",
	}

	strRecord := MapJsonToCsvString(headers, jsonData)

	assert.Equal(t, "\"val2\",\"val1\"\n", strRecord, "Json Values should be mapped to csv in header order")
}

// Tests that null values are inserted into the csv correctly
func TestMapJsonToCsvNullValue(t *testing.T) {
	headers := []string{"A", "B"}

	jsonData := map[string]interface{}{
		"A": "1",
	}

	strRecord := MapJsonToCsvString(headers, jsonData)

	assert.Equal(t, "\"1\",\n", strRecord, "Null values should be present in csv string")
}

// Tests that escapeVal wraps value in double quotes
func TestEscapeVal(t *testing.T) {
	value := "one"
	assert.Equal(t, "\"one\"", escapeVal(value))
}

// Tests that the Header array can be mapped to csv correctly
func TestHeader(t *testing.T) {
	header := []string{"one", "two", "three"}
	assert.Equal(t, "\"one\",\"two\",\"three\"\n", HeaderToCsv(header))
}
