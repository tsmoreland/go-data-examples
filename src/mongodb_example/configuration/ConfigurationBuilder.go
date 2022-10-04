package configuration

type Builder interface {
	AddJsonFile(filename string) Builder
	AddEnvironment() Builder
	AddUserSecrets() Builder
	Build() (AppConfiguration, error)
}

type appConfigurationBuilder struct {
	err               error
	hostname          string
	port              int32
	connectionOptions string
	username          string
	password          string
	secret            string
}

func NewBuilder() Builder {
	return &appConfigurationBuilder{}
}

func (b *appConfigurationBuilder) AddJsonFile(filename string) Builder {
	if b.err != nil {
		return b
	}

	config, err := newJsonConfiguration(filename)
	if err != nil {
		b.err = err
		return b
	}
	b.hostname = config.Hostname
	b.port = config.Port
	b.connectionOptions = config.ConnectionOptions
	b.username = config.Username
	b.password = config.Password
	b.secret = config.SecretId

	return b
}
func (b *appConfigurationBuilder) AddEnvironment() Builder {
	if b.err != nil {
		return b
	}

	config := newEnvironmentConfig()

	if config.Hostname != nil {
		b.hostname = *config.Hostname
	}
	if config.Port != nil {
		b.port = *config.Port
	}
	if config.ConnectionOptions != nil {
		b.connectionOptions = *config.ConnectionOptions
	}
	if config.Username != nil {
		b.username = *config.Username
	}
	if config.Password != nil {
		b.password = *config.Password
	}
	if config.SecretId != nil {
		b.secret = *config.SecretId
	}

	return b
}

func (b *appConfigurationBuilder) AddUserSecrets() Builder {
	if b.err != nil {
		return b
	}

	config, err := newSecretConfiguration(b.secret)
	if err != nil {
		// ignore error, secret is optional
		return b
	} else {
		if config.Hostname != "" {
			b.hostname = config.Hostname
		}
		if config.Port >= 0 {
			b.port = config.Port
		}
		if config.Username != "" {
			b.username = config.Username
		}
		if config.Password != "" {
			b.password = config.Password
		}
		if config.ConnectionOptions != "" {
			b.connectionOptions = config.ConnectionOptions
		}
	}

	return b
}

func (b *appConfigurationBuilder) Build() (AppConfiguration, error) {
	if b.err != nil {
		return nil, b.err
	}

	config := newConfiguration(b.hostname, b.port, b.username, b.password, b.connectionOptions)
	return config, nil
}
