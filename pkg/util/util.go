package util

import (
	"encoding/json"
	"github.com/fatih/structs"
)

func ValidateHeaders(v interface{}) bool {
	result := structs.HasZero(v)
	return !result
}
func GetHeaderMap(v interface{}, out *map[string]string) error {
	b, _ := json.Marshal(v)
	err := json.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}
