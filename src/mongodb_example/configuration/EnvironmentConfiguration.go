package configuration

import (
	"log"
	"os"
	"strconv"
)

const (
	envHostnameKey          = "GO_MYSQL_EXAMPLE__HOSTNAME"
	envPortKey              = "GO_MYSQL_EXAMPLE__PORT"
	envUsernameKey          = "GO_MYSQL_EXAMPLE__USERNAME"
	envPasswordKey          = "GO_MYSQL_EXAMPLE__PASSWORD"
	envConnectionOptionsKey = "GO_MYSQL_EXAMPLE__CONNECTION_OPTIONS"
	envSecretIdKey          = "GO_MYSQL_EXAMPLE__SECRET_ID"
)

type envConfiguration struct {
	Hostname          *string
	Port              *int32
	ConnectionOptions *string
	Username          *string
	Password          *string
	SecretId          *string
}

func newEnvironmentConfig() *envConfiguration {
	hostname := getValueOrNilFromEnv(envHostnameKey)
	portString := getValueOrNilFromEnv(envPortKey)

	var port *int32
	if portString != nil {
		portValue64, err := strconv.ParseInt(*portString, 10, 32)
		if err != nil {
			log.Printf("%v is not a valid port number: %v", portString, err)
			portValue64 = -1
		}
		portValue := int32(portValue64)
		port = &portValue
	} else {
		portValue := int32(-1)
		port = &portValue
	}

	username := getValueOrNilFromEnv(envUsernameKey)
	password := getValueOrNilFromEnv(envPasswordKey)
	connectionOptions := getValueOrNilFromEnv(envConnectionOptionsKey)
	secret := getValueOrNilFromEnv(envSecretIdKey)

	return &envConfiguration{
		Hostname:          hostname,
		Port:              port,
		Username:          username,
		Password:          password,
		ConnectionOptions: connectionOptions,
		SecretId:          secret,
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
