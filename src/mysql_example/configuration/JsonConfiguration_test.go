package configuration

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

const (
	expectedJsonAddress      = "https://example.com:6033/json"
	expectedJsonDatabaseName = "json-phi"
	expectedJsonUsername     = "json-alpha"
	expectedJsonPassword     = "json-bravo"
	expectedJsonSecretId     = "12345e2f-e70b-4201-b252-ac8668ba0dd0"
)

func Test_newJsonConfiguration_ReturnsError_WhenReadFileReturnsError(t *testing.T) {
	filename, err := arrangeTestJsonFile(false, false, t)
	if err != nil {
		t.Fatal(err)
	}

	_, err = newJsonConfiguration(filename)
	if err == nil {
		t.Fatalf("newJsonConfiguration did not return expected error")
	}
}

func Test_newJsonConfiguration_ReturnsError_WhenFileMalformed(t *testing.T) {
	filename, err := arrangeTestJsonFile(true, false, t)
	if err != nil {
		t.Fatal(err)
	}

	_, err = newJsonConfiguration(filename)
	if err == nil {
		t.Fatal("newJsonConfiguration did not return expected error")
	}
}

func Test_newJsonConfiguration_DoesNotReturnError_WhenFileIsWellFormed(t *testing.T) {
	filename, err := arrangeTestJsonFile(true, true, t)
	if err != nil {
		t.Fatal(err)
	}
	_, err = newJsonConfiguration(filename)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_newJsonConfiguration_ReturnsExpectedConfig_WhenFileIsWellFormed(t *testing.T) {
	filename, err := arrangeTestJsonFile(true, true, t)
	if err != nil {
		t.Fatal(err)
	}
	c, err := newJsonConfiguration(filename)
	if err != nil {
		t.Fatal(err)
	}

	if c.Address != expectedJsonAddress {
		t.Fatalf("address %v does not match expected value", c.Address)
	}
	if c.DatabaseName != expectedJsonDatabaseName {
		t.Fatalf("DatabaseName %v does not match expected value", c.DatabaseName)
	}
	if c.Username != expectedJsonUsername {
		t.Fatalf("username %v does not match expected value", c.Username)
	}
	if c.Password != expectedJsonPassword {
		t.Fatalf("password %v does not match expected value", c.Password)
	}
	if c.SecretId != expectedJsonSecretId {
		t.Fatalf("secretId %v does not match expected value", c.SecretId)
	}
}

func arrangeTestJsonFile(ensureExists bool, isValid bool, t *testing.T) (string, error) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "appsettings.test.json")

	_, err := os.Stat(filename)
	if ensureExists {
		return getFilepathWhenFileMustExist(filename, isValid, err)
	} else {
		return getFilepathWhenFileMustNotExist(filename, err)
	}
}

func getFilepathWhenFileMustNotExist(filename string, statError error) (string, error) {
	if errors.Is(statError, fs.ErrNotExist) {
		return filename, nil
	} else {
		err := os.Remove(filename)
		return filename, err
	}
}
func getFilepathWhenFileMustExist(filename string, isValid bool, statError error) (string, error) {
	if errors.Is(statError, fs.ErrNotExist) {
		return createTestJsonFile(filename, isValid)
	} else {
		return filename, nil // assuming it has correct format
	}
}

func createTestJsonFile(filename string, isValid bool) (string, error) {
	var content string
	if isValid {
		content = fmt.Sprintf(`{
  "address": "%v",
  "database_name": "%v",
  "username": "%v",
  "password": "%v",
  "secret": "%v"
}`, expectedJsonAddress, expectedJsonDatabaseName, expectedJsonUsername, expectedJsonPassword, expectedJsonSecretId)
	} else {
		content = fmt.Sprintf(`settings:
  address: %v
  username: %v
  password: %v
  secret: %v	
`, expectedAddress, expectedUsername, expectedPassword, expectedSecretId)
	}

	err := os.WriteFile(filename, []byte(content), 0600)
	return filename, err
}
