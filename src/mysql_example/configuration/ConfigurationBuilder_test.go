package configuration

import "testing"

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

	err = checkIfConfigurationIsExpected(c, false)
	if err != nil {
		t.Fatal(err)
	}

	if c.SecretId() != "" {
		t.Fatalf("secret was returned with configuration")
	}
}
