package configuration

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func newSecretConfiguration(secretId string) (*jsonConfiguration, error) {

	if secretId == "" {
		return nil, fmt.Errorf("secret id cannot be empty")
	}

	home, err := getHomeOrUserProfileOrError()
	if err != nil {
		return nil, err
	}

	if err = ensureHomeOrUserProfileExistsAndIsDirectory(home); err != nil {
		return nil, err
	}

	secretFile := filepath.Join(home, ".go", "secrets", secretId, "secret.json")
	if _, err = os.Stat(secretFile); err == nil {
		return newJsonConfiguration(secretFile)
	} else if errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		// error might exist, need to review err for more details which we won't do yet
		return nil, err
	}
}

func getHomeOrUserProfileOrError() (string, error) {
	home, present := os.LookupEnv("UserProfile")
	if !present {
		if home, present = os.LookupEnv("Home"); !present {
			return "", fmt.Errorf("neither userprofile nor home directory found")
		}
	}
	return home, nil
}

func ensureHomeOrUserProfileExistsAndIsDirectory(home string) error {
	fi, err := os.Stat(home)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("userprofile or home value is not a directory")
	}

	return nil
}
