package configuration

import "os"

type envConfiguration struct {
	Address  *string
	Username *string
	Password *string
	SecretId *string
}

func newEnvironmentConfig() *envConfiguration {
	address := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE__ADDRESS")
	username := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE__USERNAME")
	password := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE__PASSWORD")
	secret := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE_SECRET_ID")

	return &envConfiguration{
		Address:  address,
		Username: username,
		Password: password,
		SecretId: secret,
	}
}

func getValueOrNilFromEnv(key string) *string {
	value, present := os.LookupEnv(key)
	if present {
		return &value
	} else {
		return nil
	}
}
