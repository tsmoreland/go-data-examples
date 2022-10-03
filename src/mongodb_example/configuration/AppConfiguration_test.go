package configuration

import (
	"fmt"
	"testing"
)

const (
	expectedHostname          = "example.com"
	expectedPort              = 42
	expectedConnectionOptions = "?maxPoolSize=42&w=majority"
	expectedUsername          = "alpha"
	expectedPassword          = "bravo"
	expectedSecretId          = "68241e2f-e70b-4201-b252-ac8668ba0dd0"
)

type configPair struct {
	Name   string
	Exists bool
	Valid  bool
}

func Test_newConfigurationShouldNotReturnNil(t *testing.T) {
	c := newConfiguration(expectedHostname, expectedPort, expectedUsername, expectedPassword, expectedConnectionOptions)
	if c == nil {
		t.Fatalf("newConfiguration returned nil")
	}
}

func Test_HostnameShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedHostname, expectedPort, expectedUsername, expectedPassword, expectedConnectionOptions)
	if c.Hostname() != expectedHostname {
		t.Fatalf("Hostname %v does not equal expected value", c.Hostname())
	}
}

func Test_PortShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedHostname, expectedPort, expectedUsername, expectedPassword, expectedConnectionOptions)
	if c.Port() != expectedPort {
		t.Fatalf("Port %v does not equal expected value", c.Port())
	}
}

func Test_UsernameShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedHostname, expectedPort, expectedUsername, expectedPassword, expectedConnectionOptions)
	if c.Username() != expectedUsername {
		t.Fatalf("Username %v does not equal expected value", c.Username())
	}
}

func Test_PasswordShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedHostname, expectedPort, expectedUsername, expectedPassword, expectedConnectionOptions)
	if c.Password() != expectedPassword {
		t.Fatalf("Password %v does not equal expected value", c.Password())
	}
}

func Test_ConnectionOptionsShouldReturnValuePassedToNewConfiguration(t *testing.T) {
	c := newConfiguration(expectedHostname, expectedPort, expectedUsername, expectedPassword, expectedConnectionOptions)
	if c.ConnectionOptions() != expectedConnectionOptions {
		t.Fatalf("ConnectionOptions %v does not equal expected value", c.ConnectionOptions())
	}
}

func checkIfConfigurationMatches(
	c AppConfiguration,
	hostname string,
	port int32,
	user string,
	pw string,
	connectionOptions string,
	secret string,
	checkSecret bool) error {

	if c.Hostname() != hostname {
		return fmt.Errorf("hostname %v does not match expected value", c.Hostname())
	}
	if c.Port() != port {
		return fmt.Errorf("port %v does not match expected value", c.Port())
	}
	if c.Username() != user {
		return fmt.Errorf("username %v does not match expected value", c.Username())
	}
	if c.Password() != pw {
		return fmt.Errorf("password %v does not match expected value", c.Password())
	}
	if c.ConnectionOptions() != connectionOptions {
		return fmt.Errorf("ConnectionOptions %v does not match expected value", c.ConnectionOptions())
	}
	if checkSecret && c.SecretId() != secret {
		return fmt.Errorf("secretId %v does not match expected value", c.SecretId())
	}
	return nil

}
