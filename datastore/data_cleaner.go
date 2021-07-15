package datastore

import (
	"strings"
)

func CleanDataMap(dataMap map[string]interface{}) {

	dataMap["price"] = cleanPrice(dataMap["price"].(string))
	dataMap["bedrooms"] = cleanBedrooms(dataMap["bedrooms"].(string))
	dataMap["square_feet"] = cleanSqft(dataMap["square_feet"].(string))
}

func cleanPrice(priceStr string) string {
	priceStr = strings.ReplaceAll(priceStr, "$", "")
	priceStr = strings.ReplaceAll(priceStr, ",", "")

	return priceStr
}

func cleanBedrooms(bedroomsStr string) string {
	if strings.Contains(bedroomsStr, "Studio") {
		return "0"
	}

	return bedroomsStr
}

func cleanSqft(sqftStr string) string {
	sqftStr = strings.ReplaceAll(sqftStr, "Square", "")
	sqftStr = strings.ReplaceAll(sqftStr, "Footage", "")
	sqftStr = strings.TrimSpace(sqftStr)
	return sqftStr
}
