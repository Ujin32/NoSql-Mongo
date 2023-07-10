package fileutil

import (
	"encoding/json"
	"os"
)

func WriteToJSONFile(data interface{}) error {
	file, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(&data)
	if err != nil {
		return err
	}
	return nil
}
