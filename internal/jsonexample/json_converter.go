package jsonexample

import (
	"encoding/json"
	"fmt"
)

const valueField = "custom_body"

func ConvertFrom[E any](event E) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	b, err := json.Marshal(event)
	if err != nil {
		return result, err
	}

	result[valueField] = b
	return result, nil
}

func ConvertTo[E any](event map[string]interface{}) (E, error) {
	result := *new(E)

	data, ok := event[valueField].(string)
	if !ok {
		return result, fmt.Errorf("error convert to Event struct, %s is not exist", valueField)
	}

	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
