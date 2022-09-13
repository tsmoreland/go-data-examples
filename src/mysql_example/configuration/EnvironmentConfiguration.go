package configuration

import "os"

const (
	envAddressKey  = "GO_MYSQL_EXAMPLE__ADDRESS"
	envUsernameKey = "GO_MYSQL_EXAMPLE__USERNAME"
	envPasswordKey = "GO_MYSQL_EXAMPLE__PASSWORD"
	envSecretIdKey = "GO_MYSQL_EXAMPLE__SECRET_ID"
)

type envConfiguration struct {
	Address  *string
	Username *string
	Password *string
	SecretId *string
}

func newEnvironmentConfig() *envConfiguration {
	address := getValueOrNilFromEnv(envAddressKey)
	username := getValueOrNilFromEnv(envUsernameKey)
	password := getValueOrNilFromEnv(envPasswordKey)
	secret := getValueOrNilFromEnv(envSecretIdKey)

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
