package configuration

import (
	"strconv"
	"testing"
)

func Test_AddJsonFileShouldSetErrorWhenFileNotFound(t *testing.T) {
	filename, err := arrangeTestJsonFile(false, false, t)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewBuilder().
		AddJsonFile(filename).
		Build()
	if err == nil {
		t.Fatalf("error is nil, expected failure due to file not found")
	}
}

func Test_AddJsonShouldNotSetErrorWhenFileExistsAndIsValid(t *testing.T) {
	filename, err := arrangeTestJsonFile(true, true, t)
	if err != nil {
		t.Fatal(err)
	}

	c, err := NewBuilder().AddJsonFile(filename).Build()

	if err != nil {
		t.Fatal(err)
	}

	if c == nil {
		t.Fatalf("configuration is nil")
	}

	err = checkIfConfigurationMatches(c,
		expectedJsonHostname,
		expectedJsonPort,
		expectedJsonUsername,
		expectedJsonPassword,
		expectedJsonConnectionOptions,
		expectedJsonSecretId,
		false)
	if err != nil {
		t.Fatal(err)
	}

	if c.SecretId() != "" {
		t.Fatalf("secret was returned with configuration")
	}
}

func Test_AddEnvironmentShouldSetNotSetErrorWhenNotFound(t *testing.T) {
	_, err := NewBuilder().AddEnvironment().Build()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_AddEnvironmentShouldSetNotSetErrorWhenFound(t *testing.T) {
	arrangeEnvSettings(t)
	_, err := NewBuilder().AddEnvironment().Build()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_AddEnvironmentShouldReturnExpectedEnvWhenPresentInEnv(t *testing.T) {
	arrangeEnvSettings(t)
	c, err := NewBuilder().AddEnvironment().Build()
	if err != nil {
		t.Fatal(err)
	}
	err = checkIfConfigurationMatches(c,
		expectedEnvHostname,
		expectedEnvPort,
		expectedEnvUsername,
		expectedEnvPassword,
		expectedEnvConnectionOptions,
		expectedEnvSecretId,
		false)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_AddUserSecretsShouldNotReturnErrorWhen(t *testing.T) {

	configPairs := []configPair{
		{"fileNotFound", false, false},
		{"fleFoundButInvalid", true, false},
		{"validFileFound", true, true},
	}

	for _, pair := range configPairs {
		t.Run(pair.Name, func(t *testing.T) {
			_, _, err := arrangeSecretFile(t, pair.Exists, pair.Valid)
			if err != nil {
				t.Fatal(err)
			}
			_, err = NewBuilder().AddUserSecrets().Build()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func Test_BuildReturnsShouldReturnValuesFromLastAddedSource(t *testing.T) {
	filename, err := arrangeTestJsonFile(true, true, t)
	if err != nil {
		t.Fatal(err)
	}
	arrangeEnvSettings(t)
	c, err := NewBuilder().
		AddJsonFile(filename).
		AddEnvironment().
		Build()

	if err != nil {
		t.Fatal(err)
	}

	err = checkIfConfigurationMatches(c,
		expectedEnvHostname,
		expectedEnvPort,
		expectedEnvUsername,
		expectedEnvPassword,
		expectedEnvConnectionOptions,
		expectedEnvSecretId,
		false)

	if err != nil {
		t.Fatal(err)
	}
}

func arrangeEnvSettings(t *testing.T) {
	t.Setenv(envHostnameKey, expectedEnvHostname)
	t.Setenv(envPortKey, strconv.Itoa(expectedEnvPort))
	t.Setenv(envUsernameKey, expectedEnvUsername)
	t.Setenv(envPasswordKey, expectedEnvPassword)
	t.Setenv(envConnectionOptionsKey, expectedEnvConnectionOptions)
	t.Setenv(envSecretIdKey, expectedEnvSecretId)
}
