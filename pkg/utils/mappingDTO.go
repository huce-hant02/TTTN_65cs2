package utils

import "encoding/json"

func MappingInterface(input interface{}, output interface{}) error {
	// Mapping input to output
	inputJson, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(inputJson, output)
	if err != nil {
		return err
	}
	return nil
}
