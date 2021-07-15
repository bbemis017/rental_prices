package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that CleanDataMap cleans a map correctly
func TestCleanDataMap(t *testing.T) {
	data := make(map[string]interface{})

	data["price"] = "$1,000"
	data["bedrooms"] = "Studio"
	data["square_feet"] = "632 Square Footage"

	CleanDataMap(data)

	assert.Equal(t, "1000", data["price"])
	assert.Equal(t, "0", data["bedrooms"])
	assert.Equal(t, "632", data["square_feet"])
}

func TestCleanPriceStripDollarSign(t *testing.T) {
	assert.Equal(t, "100", cleanPrice("$100"))
}

func TestCleanPriceStripComma(t *testing.T) {
	assert.Equal(t, "1000", cleanPrice("1,000"))
}

func TestCleanBedroomsConvertStudio(t *testing.T) {
	assert.Equal(t, "0", cleanBedrooms("Studio"))
}

func TestCleanSquareFeet(t *testing.T) {
	assert.Equal(t, "632", cleanSqft("632 Square Footage"))
}
