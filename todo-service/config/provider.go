package config

type Configs struct {
	configs map[string]interface{}
}

type Provider interface {
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetFloat(key string) (float64, error)
	loadConfig() error
	logConfigs()
}
