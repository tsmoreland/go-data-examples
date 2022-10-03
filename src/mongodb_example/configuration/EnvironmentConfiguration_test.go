package configuration

import (
	"os"
	"strconv"
	"testing"
)

const (
	expectedEnvHostname          = "example.com"
	expectedEnvPort              = 29019
	expectedEnvUsername          = "env-alpha"
	expectedEnvPassword          = "env-bravo"
	expectedEnvConnectionOptions = "env=phi"
	expectedEnvSecretId          = "22346e2f-e70b-4201-b252-ac8668ba0dd0"
)

func Test_newEnvironmentConfigShouldReturnNilHostnameWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envHostnameKey); present {
		t.Skipf("Hostnames already defined")
	}
	c := newEnvironmentConfig()
	if c.Hostname != nil {
		t.Fatalf("Hostname is not nil")
	}
}
func Test_newEnvironmentConfigShouldReturnNilPortWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envPortKey); present {
		t.Skipf("Ports already defined")
	}
	c := newEnvironmentConfig()
	if c.Port != nil {
		t.Fatalf("Port is not nil")
	}
}

func Test_newEnvironmentConfigShouldReturnNilUsernameWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envUsernameKey); present {
		t.Skipf("Usernames already defined")
	}
	c := newEnvironmentConfig()
	if c.Username != nil {
		t.Fatalf("Username is not nil")
	}
}

func Test_newEnvironmentConfigShouldReturnNilPasswordWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envPasswordKey); present {
		t.Skipf("Passwords already defined")
	}
	c := newEnvironmentConfig()
	if c.Password != nil {
		t.Fatalf("Password is not nil")
	}
}

func Test_newEnvironmentConfigShouldReturnNilConnectionOptionsWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envConnectionOptionsKey); present {
		t.Skipf("Passwords already defined")
	}
	c := newEnvironmentConfig()
	if c.ConnectionOptions != nil {
		t.Fatalf("Password is not nil")
	}
}

func Test_newEnvironmentConfigShouldReturnNilSecretIdWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envSecretIdKey); present {
		t.Skipf("SecretIds already defined")
	}
	c := newEnvironmentConfig()
	if c.SecretId != nil {
		t.Fatalf("SecretId is not nil")
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValueHostnameWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envHostnameKey, expectedEnvHostname)

	c := newEnvironmentConfig()
	if c.Hostname == nil {
		t.Fatalf("Hostname not found")
	}
	if *c.Hostname != expectedEnvHostname {
		t.Fatalf("Hostname %v does not match expected value", c.Hostname)
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValuePortWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envPortKey, strconv.Itoa(expectedEnvPort))

	c := newEnvironmentConfig()
	if c.Port == nil {
		t.Fatalf("Port not found")
	}
	if *c.Port != expectedEnvPort {
		t.Fatalf("Port %v does not match expected value", c.Port)
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValueConnectionOptionsWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envConnectionOptionsKey, expectedEnvConnectionOptions)

	c := newEnvironmentConfig()
	if c.ConnectionOptions == nil {
		t.Fatalf("ConnectionOptions not found")
	}
	if *c.ConnectionOptions != expectedEnvConnectionOptions {
		t.Fatalf("ConnectionOptions %v does not match expected value", *c.ConnectionOptions)
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValueUsernameWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envUsernameKey, expectedEnvUsername)

	c := newEnvironmentConfig()
	if c.Username == nil {
		t.Fatalf("Username not found")
	}
	if *c.Username != expectedEnvUsername {
		t.Fatalf("Username %v does not match expected value", c.Username)
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValuePasswordWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envPasswordKey, expectedEnvPassword)
	c := newEnvironmentConfig()
	if c.Password == nil {
		t.Fatalf("Password not found")
	}
	if *c.Password != expectedEnvPassword {
		t.Fatalf("Password %v does not match expected value", c.Password)
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValueSecretIdWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envSecretIdKey, expectedEnvSecretId)
	c := newEnvironmentConfig()
	if c.SecretId == nil {
		t.Fatalf("SecretId not found")
	}
	if *c.SecretId != expectedEnvSecretId {
		t.Fatalf("SecretId %v does not match expected value", c.SecretId)
	}
}
