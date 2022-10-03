package configuration

import (
	"os"
	"testing"
)

const (
	expectedEnvAddress      = "https://example.com:6033/env"
	expectedEnvDatabaseName = "env-phi"
	expectedEnvUsername     = "env-alpha"
	expectedEnvPassword     = "env-bravo"
	expectedEnvSecretId     = "22346e2f-e70b-4201-b252-ac8668ba0dd0"
)

func Test_newEnvironmentConfigShouldReturnNilAddressWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envAddressKey); present {
		t.Skipf("Addresss already defined")
	}
	c := newEnvironmentConfig()
	if c.Address != nil {
		t.Fatalf("Address is not nil")
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

func Test_newEnvironmentConfigShouldReturnNilDatabaseNameWhenEnvironmentVariableNotSet(t *testing.T) {
	if _, present := os.LookupEnv(envDatabaseNameKey); present {
		t.Skipf("Passwords already defined")
	}
	c := newEnvironmentConfig()
	if c.DatabaseName != nil {
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

func Test_newEnvironmentConfigShouldReturnExpectedValueAddressWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envAddressKey, expectedEnvAddress)

	c := newEnvironmentConfig()
	if c.Address == nil {
		t.Fatalf("Address not found")
	}
	if *c.Address != expectedEnvAddress {
		t.Fatalf("Address %v does not match expected value", c.Address)
	}
}

func Test_newEnvironmentConfigShouldReturnExpectedValueDatabaseNameWhenEnvironmentVariableSet(t *testing.T) {
	t.Setenv(envDatabaseNameKey, expectedEnvDatabaseName)

	c := newEnvironmentConfig()
	if c.DatabaseName == nil {
		t.Fatalf("DatabaseName not found")
	}
	if *c.DatabaseName != expectedEnvDatabaseName {
		t.Fatalf("DatabaseName %v does not match expected value", *c.DatabaseName)
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
