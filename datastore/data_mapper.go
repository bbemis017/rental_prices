package datastore

// mapJsonToCsvString extracts values from json map and returns string in quoted csv format
//  header - array of json key names to extract in order
//  json   - map of string key/value pairs from json
func MapJsonToCsvString(header []string, json map[string]interface{}) string {

	csvRecord := ""
	for index, col := range header {

		// if a value does not exist in the json it should be left blank in the csv but the delimiters should still be inserted
		// For Example:
		// A,,B
		if json[col] != nil {
			csvRecord += escapeVal(json[col].(string))
		}

		if index < len(header)-1 {
			csvRecord += ","
		}

	}

	return csvRecord + "\n"
}

func HeaderToCsv(header []string) string {
	record := ""
	for index, col := range header {
		record += escapeVal(col)
		if index < len(header)-1 {
			record += ","
		}
	}
	return record + "\n"
}

// escapeVal wraps string value in double quotes
// Example:
// abc -> "abc"
func escapeVal(val string) string {
	return "\"" + val + "\""
}
