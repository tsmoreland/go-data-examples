//go:build !windows

package configuration

func getSecretHome(secretId string) (string, error) {
	if secretId == "" {
		return "", fmt.Errorf("secret id cannot be empty")
	}

	home, present := os.LookupEnv("Home")
	if present {
		return home, nil
	} else {
		return "", fmt.Errorf("user profile not found")
	}
}
