package configuration

type AppConfiguration interface {
	Hostname() string
	Port() int32
	Username() string
	Password() string
	ConnectionOptions() string
	SecretId() string
}

type readonlyConfiguration struct {
	hostname          string
	port              int32
	username          string
	password          string
	connectionOptions string
	secretId          string
}

func newConfiguration(
	hostname string,
	port int32,
	username string,
	password string,
	connectionOptions string) *readonlyConfiguration {

	return &readonlyConfiguration{
		hostname:          hostname,
		port:              port,
		connectionOptions: connectionOptions,
		username:          username,
		password:          password,
		secretId:          "",
	}
}

func (c readonlyConfiguration) Hostname() string {
	return c.hostname
}

func (c readonlyConfiguration) Port() int32 {
	return c.port
}

func (c readonlyConfiguration) Username() string {
	return c.username
}

func (c readonlyConfiguration) Password() string {
	return c.password
}

func (c readonlyConfiguration) ConnectionOptions() string {
	return c.connectionOptions
}

func (c readonlyConfiguration) SecretId() string {
	return c.secretId
}
