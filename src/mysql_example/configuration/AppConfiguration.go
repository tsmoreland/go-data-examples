package configuration

type AppConfiguration interface {
	Address() string
	DatabaseName() string
	Username() string
	Password() string
	SecretId() string
}

type readonlyConfiguration struct {
	address      string
	username     string
	password     string
	databaseName string
	secretId     string
}

func newConfiguration(
	address string,
	databaseName string,
	username string,
	password string) *readonlyConfiguration {

	return &readonlyConfiguration{
		address:      address,
		databaseName: databaseName,
		username:     username,
		password:     password,
		secretId:     "",
	}
}

func (c readonlyConfiguration) Address() string {
	return c.address
}

func (c readonlyConfiguration) DatabaseName() string {
	return c.databaseName
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
