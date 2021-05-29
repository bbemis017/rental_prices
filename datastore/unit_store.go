package datastore

import (
	"log"
	"strconv"
	"strings"
)

type Unit struct {
	UnitNumber   string
	Price        int
	Availability string
	Bedrooms     int
	Baths        int
	Address      string
	Complex      string
	CreatedAt    string
}

func NewUnit(rawMap map[string]interface{}, complex string, address string, createdAt string) (Unit, error) {
	return Unit{
		Complex:      complex,
		Address:      address,
		UnitNumber:   rawMap["Unit"].(string),
		Price:        convertStringPriceToInt(rawMap["Price"].(string)),
		Availability: rawMap["Availability"].(string),
		Bedrooms:     convertStringBedroomsToInt(rawMap["Bedrooms"].(string)),
		Baths:        convertStringToInt(rawMap["Baths"].(string)),
		CreatedAt:    createdAt,
	}, nil
}

func (unit *Unit) Save(store *CSVStore) {
	var record []string

	record = append(record, unit.CreatedAt)
	record = append(record, unit.Complex)
	record = append(record, unit.UnitNumber)
	record = append(record, strconv.Itoa(unit.Price))
	record = append(record, unit.Availability)
	record = append(record, strconv.Itoa(unit.Bedrooms))
	record = append(record, strconv.Itoa(unit.Baths))
	record = append(record, unit.Address)

	store.WriteRecord(record)
}

func WriteHeader(store *CSVStore) {
	header := []string{"created_at", "complex", "unit_number", "price", "availability", "bedrooms", "baths", "address"}
	store.WriteRecord(header)
}

func convertStringPriceToInt(priceStr string) int {
	priceStr = strings.ReplaceAll(priceStr, "$", "")
	priceStr = strings.ReplaceAll(priceStr, ",", "")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		log.Fatalf("Error converting Price %s to Integer", priceStr)
	}
	return price
}

func convertStringBedroomsToInt(bedroomsStr string) int {
	if strings.Contains(bedroomsStr, "Studio") {
		return 0
	}

	bedrooms, err := strconv.Atoi(bedroomsStr)
	if err != nil {
		log.Fatalf("Error converting Bedrooms %s to Integer", bedroomsStr)
	}
	return bedrooms
}

func convertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("String %s is not an integer", str)
	}
	return i
}
