package datastore

import (
	"strings"
)

func CleanDataMap(dataMap map[string]interface{}) {

	dataMap["price"] = cleanPrice(dataMap["price"].(string))
	dataMap["bedrooms"] = cleanBedrooms(dataMap["bedrooms"].(string))
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
