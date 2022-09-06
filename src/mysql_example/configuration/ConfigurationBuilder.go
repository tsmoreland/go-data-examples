package configuration

type Builder interface {
	AddJsonFile(filename string) Builder
	AddEnvironment() Builder
	AddUserSecrets() Builder
	Build() (AppConfiguration, error)
}

type appConfigurationBuilder struct {
	err      error
	address  string
	username string
	password string
}

func NewBuilder() Builder {
	return &appConfigurationBuilder{}
}

func (b *appConfigurationBuilder) AddJsonFile(filename string) Builder {

	return b
}
func (b *appConfigurationBuilder) AddEnvironment() Builder {
	return b
}

func (b *appConfigurationBuilder) AddUserSecrets() Builder {
	return b
}

func (b *appConfigurationBuilder) Build() (AppConfiguration, error) {
	if b.err != nil {
		return nil, b.err
	}

	config := newConfiguration(b.address, b.username, b.password)
	return config, nil
}
