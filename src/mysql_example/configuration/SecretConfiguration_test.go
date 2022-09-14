package configuration

import "testing"

func Test_newSecretConfigurationShouldReturnErrorWhenFileNotFound(t *testing.T) {
	home := t.TempDir()
	t.Setenv(secretHomeEnvVariable, home)

	_, err := newSecretConfiguration(expectedSecretId)
	if err == nil {
		t.Fatalf("newSecretConfiguration did not return expected error")
	}
}
