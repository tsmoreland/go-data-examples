package configuration

import (
	"fmt"
	"os"
)

func getSecretHome(secretId string) (string, error) {
	if secretId == "" {
		return "", fmt.Errorf("secret id cannot be empty")
	}

	home, present := os.LookupEnv("UserProfile")
	if present {
		return home, nil
	} else {
		return "", fmt.Errorf("user profile not found")
	}
}
