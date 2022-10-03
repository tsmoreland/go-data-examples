package configuration

import (
	"fmt"
	"testing"
)

const (
	expectedAddress      = "https://example.com:6033"
	expectedDatabaseName = "phi"
	expectedUsername     = "alpha"
	expectedPassword     = "bravo"
	expectedSecretId     = "68241e2f-e70b-4201-b252-ac8668ba0dd0"
)

type configPair struct {
	Name   string
	Exists bool
	Valid  bool
}

func Test_newConfigurationShouldNotReturnNil(t *testing.T) {
	c := newConfiguration(expectedAddress, expectedDatabaseName, expectedUsername, expectedPassword)
	if c == nil {
		t.Fatalf("newConfiguration returned nil")
	}
}

func Test_AddressShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedAddress, expectedDatabaseName, expectedUsername, expectedPassword)
	if c.Address() != expectedAddress {
		t.Fatalf("Address %v does not equal expected value", c.Address())
	}
}

func Test_DatabaseNameShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedAddress, expectedDatabaseName, expectedUsername, expectedPassword)
	if c.DatabaseName() != expectedDatabaseName {
		t.Fatalf("DatabaseName %v does not equal expected value", c.DatabaseName())
	}
}

func Test_UsernameShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedAddress, expectedDatabaseName, expectedUsername, expectedPassword)
	if c.Username() != expectedUsername {
		t.Fatalf("Username %v does not equal expected value", c.Username())
	}
}

func Test_PasswordShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedAddress, expectedDatabaseName, expectedUsername, expectedPassword)
	if c.Password() != expectedPassword {
		t.Fatalf("Password %v does not equal expected value", c.Password())
	}
}

func checkIfConfigurationMatches(
	c AppConfiguration,
	addr string,
	databaseName string,
	user string,
	pw string,
	secret string,
	checkSecret bool) error {

	if c.Address() != addr {
		return fmt.Errorf("address %v does not match expected value", c.Address())
	}
	if c.DatabaseName() != databaseName {
		return fmt.Errorf("DatabaseName %v does not match expected value", c.DatabaseName())
	}
	if c.Username() != user {
		return fmt.Errorf("username %v does not match expected value", c.Username())
	}
	if c.Password() != pw {
		return fmt.Errorf("password %v does not match expected value", c.Password())
	}
	if checkSecret && c.SecretId() != secret {
		return fmt.Errorf("secretId %v does not match expected value", c.SecretId())
	}
	return nil

}
