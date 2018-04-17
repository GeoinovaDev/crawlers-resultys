package convert

// ArrayInterfaceToArrayString convert array interface to array string
// Return array de string
func ArrayInterfaceToArrayString(arr []interface{}) []string {
	result := []string{}

	for i := 0; i < len(arr); i++ {
		result = append(result, arr[i].(string))
	}

	return result
}
