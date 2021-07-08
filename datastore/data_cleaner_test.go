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

	CleanDataMap(data)

	assert.Equal(t, "1000", data["price"])
	assert.Equal(t, "0", data["bedrooms"])
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
