package main

import (
	"fmt"
	"github.com/tsmoreland/go-data-examples/mongodb_example/configuration"
	"strings"
)

func buildConnectionString(c configuration.AppConfiguration) (string, error) {
	if c == nil {
		return "", fmt.Errorf("invalid argument")
	}
	hostname := c.Hostname()
	port := c.Port()
	username := c.Username()
	password := c.Password()
	connectionOptions := c.ConnectionOptions()

	if err := isEmptyOrWhiteSpace(hostname, "hostname"); err != nil {
		return "", err
	}
	if port < 0 {
		port = 27017
	}

	usernameIsEmpty := false
	passwordIsEmpty := false

	if err := isEmptyOrWhiteSpace(username, "username"); err != nil {
		usernameIsEmpty = true
	}
	if err := isEmptyOrWhiteSpace(password, "password"); err != nil {
		passwordIsEmpty = true
	}

	var authString string
	if usernameIsEmpty && passwordIsEmpty {
		authString = ""
	} else if !usernameIsEmpty && !passwordIsEmpty {
		authString = fmt.Sprintf("%v:%v@", username, password)
	} else {
		return "", fmt.Errorf("username or password is empty, either both or neither must have a value")
	}

	if err := isEmptyOrWhiteSpace(connectionOptions, ""); err != nil {
		conn := fmt.Sprintf("mongodb://%v%v:%v/", authString, hostname, port)
		return conn, nil
	} else {
		conn := fmt.Sprintf("mongodb://%v%v:%v/?%v", authString, hostname, port, connectionOptions)
		return conn, nil
	}
}

func isEmptyOrWhiteSpace(s string, name string) error {
	if len(strings.Trim(s, " \t\r\n")) == 0 {
		return fmt.Errorf("invalid %v", name)
	} else {
		return nil
	}

}
