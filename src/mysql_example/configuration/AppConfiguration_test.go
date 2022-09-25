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

func checkIfConfigurationIsExpected(c AppConfiguration, checkSecret bool) error {
	if c == nil {
		return fmt.Errorf("configuration is nil")
	}
	if c.Address() != expectedAddress {
		return fmt.Errorf("address %v does not match expected value", c.Address())
	}
	if c.DatabaseName() != expectedDatabaseName {
		return fmt.Errorf("DatabaseName %v does not match expected value", c.DatabaseName())
	}
	if c.Username() != expectedUsername {
		return fmt.Errorf("username %v does not match expected value", c.Username())
	}
	if c.Password() != expectedPassword {
		return fmt.Errorf("password %v does not match expected value", c.Password())
	}
	if checkSecret && c.SecretId() != expectedSecretId {
		return fmt.Errorf("secretId %v does not match expected value", c.SecretId())
	}
	return nil
}
