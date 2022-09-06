package configuration

import (
	"encoding/json"
	"os"
)

type jsonConfiguration struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	SecretId string `json:"secret"`
}

func newJsonConfiguration(filename string) (*jsonConfiguration, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config jsonConfiguration
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
