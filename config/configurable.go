package config

type Configurable interface {
	Configure(config interface{})
	GetConfigPath() string
	SetConfigPath(path string)
}
