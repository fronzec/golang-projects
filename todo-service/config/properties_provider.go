package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type PropertiesConfigProvider struct {
	Configs
}

func (provider *PropertiesConfigProvider) logConfigs() {
	fmt.Printf("=================================================\n")
	for key, value := range provider.configs {
		fmt.Printf("config: <%v=%v>\n", key, value)
	}
	fmt.Printf("=================================================\n")
}

func NewPropertiesConfigProvider() (Provider, error) {
	provider := PropertiesConfigProvider{}
	err := provider.loadConfig()
	provider.logConfigs()
	if err != nil {
		return nil, err
	}
	return &provider, err
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
	configs := map[string]interface{}{}
	fileName, ok := os.LookupEnv("PROPERTIES_FILENAME")
	if !ok {
		fmt.Printf("cannot read PROPERTIES_FILENAME environment, default will be used\n")
		fileName = "./config/example/app.properties"
	}
	fmt.Printf("reading properfies file:%v\n", fileName)
	open, err2 := os.Open(fileName)
	if err2 != nil {
		message := fmt.Sprintf("cannot open properties file %v error:%v\n", fileName, err2)
		fmt.Printf(message)
		return errors.New(message)
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			fmt.Printf("cannot close file:%v\n", open)
		}
	}(open)

	scanner := bufio.NewScanner(open)

	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal > 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				configs[key] = value
			}
		}
	}

	if err3 := scanner.Err(); err3 != nil {
		fmt.Printf("something goes wrong with the scanner %v\n", err3)
		return err3
	}
	provider.configs = configs
	return nil
}
