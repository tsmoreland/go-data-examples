package configuration

type AppConfiguration interface {
	Address() string
	Username() string
	Password() string
	SecretId() string
}

type readonlyConfiguration struct {
	address  string
	username string
	password string
	secretId string
}

func newConfiguration(
	address string,
	username string,
	password string) *readonlyConfiguration {

	return &readonlyConfiguration{
		address:  address,
		username: username,
		password: password,
		secretId: "",
	}
}

func (c readonlyConfiguration) Address() string {
	return c.address
}

func (c readonlyConfiguration) Username() string {
	return c.username
}

func (c readonlyConfiguration) Password() string {
	return c.password
}

func (c readonlyConfiguration) SecretId() string {
	return c.secretId
}
