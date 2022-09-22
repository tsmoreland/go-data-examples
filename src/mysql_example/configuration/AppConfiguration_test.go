package configuration

import "testing"

func Test_newConfigurationShouldNotReturnNil(t *testing.T) {
	c := newConfiguration(expectedAddress, expectedDatabaseName, expectedUsername, expectedPassword)
	if c == nil {
		t.Fatalf("newConfiguration returned nil")
	}

}
