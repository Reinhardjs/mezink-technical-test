package repository

import (
	"encoding/json"
)

// StringToInts converts a comma-separated string to a slice of ints
func StringToInts(s string) *[]int {
	var result []int
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil
	}
	return &result
}
