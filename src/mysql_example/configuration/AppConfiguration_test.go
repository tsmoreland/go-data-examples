package configuration

import "testing"

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
