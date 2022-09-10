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
	expectedAddress  = "https://example.com:6033"
	expectedUsername = "alpha"
	expectedPassword = "bravo"
	expectedSecretId = "68241e2f-e70b-4201-b252-ac8668ba0dd0"
)

func Test_newJsonConfiguration_ReturnsError_WhenReadFileReturnsError(t *testing.T) {
	filename, err := arrangeFile(false, t)
	if err != nil {
		t.Fatal(err)
	}

	_, err = newJsonConfiguration(filename)
	if err == nil {
		t.Fatalf("newJsonConfiguration did not return expected error")
	}
}

func arrangeFile(ensureExists bool, t *testing.T) (string, error) {

	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "appsettings.test.json")

	_, err := os.Stat(filename)
	if ensureExists {
		return getFilepathWhenFileMustExist(filename, err)
	} else {
		return getFilepathWhenFileMustNotExist(filename, err)
	}
}

func getFilepathWhenFileMustExist(filename string, statError error) (string, error) {
	if errors.Is(statError, fs.ErrNotExist) {
		return filename, nil
	} else {
		err := os.Remove(filename)
		return filename, err
	}
}
func getFilepathWhenFileMustNotExist(filename string, statError error) (string, error) {
	if errors.Is(statError, fs.ErrNotExist) {
		return createFile(filename)
	} else {
		return filename, nil // assuming it has correct format
	}
}

func createFile(filename string) (string, error) {
	json := fmt.Sprintf(`{
  "address": "%v",
  "username": "%v",
  "password": "%v",
  "secretId": "%v",
}`, expectedAddress, expectedUsername, expectedPassword, expectedSecretId)

	err := os.WriteFile(filename, []byte(json), 0600)
	return filename, err
}
