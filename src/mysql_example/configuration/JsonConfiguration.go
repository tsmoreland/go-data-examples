package configuration

type JsonConfiguration struct {
	filename string
}

func newJsonConfiguration(filename string) *JsonConfiguration {
	return &JsonConfiguration{filename: filename}
}
