package configuration

import "os"

type envConfiguration struct {
	Address  *string
	Username *string
	Password *string
}

func newEnvironmentConfig() *envConfiguration {
	address := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE__ADDRESS")
	username := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE__USERNAME")
	password := getValueOrNilFromEnv("GO_MYSQL_EXAMPLE__PASSWORD")

	return &envConfiguration{
		Address:  address,
		Username: username,
		Password: password,
	}
}

func getValueOrNilFromEnv(key string) *string {
	value, present := os.LookupEnv("GO_MYSQL_EXAMPLE__ADDRESS")
	if present {
		return &value
	} else {
		return nil
	}
}
