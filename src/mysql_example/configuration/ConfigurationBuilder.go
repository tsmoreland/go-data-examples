package configuration

type Builder interface {
	AddJsonFile(filename string) Builder
	AddEnvironment() Builder
	AddUserSecrets() Builder
	Build() (AppConfiguration, error)
}

type appConfigurationBuilder struct {
	err          error
	address      string
	databaseName string
	username     string
	password     string
	secret       string
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
	}
	b.address = config.Address
	b.databaseName = config.DatabaseName
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

	if config.Address != nil {
		b.address = *config.Address
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
		if config.Address != "" {
			b.address = config.Address
		}
		if config.Username != "" {
			b.username = config.Username
		}
		if config.Password != "" {
			b.password = config.Password
		}
	}

	return b
}

func (b *appConfigurationBuilder) Build() (AppConfiguration, error) {
	if b.err != nil {
		return nil, b.err
	}

	config := newConfiguration(b.address, b.databaseName, b.username, b.password)
	return config, nil
}
