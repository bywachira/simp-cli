package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ConfigurationPayload file to get the
type ConfigurationPayload struct {
	Commands []string `json:"commands"`
}

// ParseJSONFile reads config file
func ParseJSONFile() ConfigurationPayload {
	jsonFile, err := os.Open("./simp.json")

	if err != nil {
		log.Fatal("`simp.json` file not found")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var configuration ConfigurationPayload

	json.Unmarshal(byteValue, &configuration)

	return configuration
}

// LogConfiguration logs config
func LogConfiguration(config ConfigurationPayload) {
	fmt.Println(config)
}
