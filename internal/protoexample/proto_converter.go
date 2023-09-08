package protoexample

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

const valueField = "custom_proto_body"

func ConvertFrom(event *Event) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	b, err := proto.Marshal(event)
	if err != nil {
		return result, err
	}

	result[valueField] = b
	return result, nil
}

func ConvertTo(event map[string]interface{}) (*Event, error) {
	result := Event{}

	data, ok := event[valueField].(string)
	if !ok {
		return &result, fmt.Errorf("error convert to Event struct, %s is not exist", valueField)
	}

	err := proto.Unmarshal([]byte(data), &result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
