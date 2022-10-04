package configuration

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	expectedSecretHostname          = "example.com"
	expectedSecretPort              = 28142
	expectedSecretConnectionOptions = "secret=phi"
	expectedSecretUsername          = "secret-alpha"
	expectedSecretPassword          = "secret-bravo"
)

func Test_newSecretConfigurationShouldReturnErrorWhenFileNotFound(t *testing.T) {
	if _, _, err := arrangeSecretFile(t, false, false); err != nil {
		t.Fatal(err)
	}

	_, err := newSecretConfiguration(expectedSecretId)
	if err == nil {
		t.Fatalf("newSecretConfiguration did not return expected error")
	}
}

func Test_newSecretConfigurationShouldReturnErrorWhenFileExistsButIsInvalid(t *testing.T) {
	if _, _, err := arrangeSecretFile(t, true, false); err != nil {
		t.Fatal(err)
	}
	_, err := newSecretConfiguration(expectedSecretId)
	if err == nil {
		t.Fatalf("newSecretConfiguration did not return expected error")
	}
}

func Test_newSecretConfigurationShouldNotReturnErrorWhenFileExistsAndValid(t *testing.T) {
	if _, _, err := arrangeSecretFile(t, true, true); err != nil {
		t.Fatal(err)
	}

	c, err := newSecretConfiguration(expectedSecretId)
	if err != nil {
		t.Fatal(err)
	}

	if c.Hostname != expectedSecretHostname {
		t.Fatalf("Hostname %v does not match expected value", c.Hostname)
	}

	if c.Port != expectedSecretPort {
		t.Fatalf("Port %v does not match expected value", c.Port)
	}

	if c.ConnectionOptions != expectedSecretConnectionOptions {
		t.Fatalf("ConnectionOptions %v does not match expected value", c.ConnectionOptions)
	}

	if c.Username != expectedSecretUsername {
		t.Fatalf("Username %v does not match expected value", c.Username)
	}

	if c.Password != expectedSecretPassword {
		t.Fatalf("Password %v does not match expected value", c.Password)
	}

}

func arrangeSecretFile(t *testing.T, exists bool, isValid bool) (string, string, error) {
	home := t.TempDir()
	t.Setenv(secretHomeEnvVariable, home)
	home = filepath.Join(home, ".go", "secrets", expectedSecretId)
	filename := filepath.Join(home, "secret.json")

	_, err := os.Stat(filename)
	if !exists {
		if os.IsNotExist(err) {
			return home, filename, nil
		} else {
			err = os.Remove(filename)
			return home, filename, err
		}
	}

	fi, err := os.Stat(home)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0700); err != nil {
			t.Fatal(err)
		}
	} else if !fi.IsDir() {
		t.Fatalf("%v is not a directory", home)
	}

	return home, filename, createTestSecretFile(filename, isValid)
}

func createTestSecretFile(filename string, isValid bool) error {
	var content string
	if isValid {
		content = fmt.Sprintf(`{
  "hostname": "%v",
  "port": "%v",
  "username": "%v",
  "password": "%v"
  "connection_options": "%v",
}`, expectedSecretHostname, expectedSecretPort, expectedSecretUsername, expectedSecretPassword,
			expectedSecretConnectionOptions)
	} else {
		content = fmt.Sprintf(`settings:
  hostname: %v
  username: %v
  password: %v
`, expectedHostname, expectedUsername, expectedPassword)
	}

	err := os.WriteFile(filename, []byte(content), 0600)
	return err
}
