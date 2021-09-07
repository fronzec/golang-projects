package config

type PropertiesConfigProvider struct {
	configs map[string]interface{}
}

func NewPropertiesConfigProvider() (*PropertiesConfigProvider,error) {
	provider := PropertiesConfigProvider{}
	err := provider.loadConfig()
	if err != nil {
		return nil, err
	}
	return &provider,err
}

func (provider *PropertiesConfigProvider) GetString(key string) (string, error) {
	panic("implement me")
}

func (provider *PropertiesConfigProvider) GetInt(key string) (int, error) {
	panic("implement me")
}

func (provider *PropertiesConfigProvider) GetInt64(key string) (int64, error) {
	panic("implement me")
}

func (provider *PropertiesConfigProvider) GetFloat(key string) (float64, error) {
	panic("implement me")
}

func (provider *PropertiesConfigProvider) loadConfig() error {
	panic("implement me")
}

