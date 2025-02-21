package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	envVars map[string]string
}

var requiredEnvVars = []string{"REDIS_ADDR"}

func NewConfig() *Config {
	envVars := make(map[string]string)

	file, err := os.Open("config.properties")
	if err != nil {
		log.Fatalf("Failed to open config.properties file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			envVars[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading config.properties file: %v", err)
	}

	for _, key := range requiredEnvVars {
		if _, exists := envVars[key]; !exists {
			log.Fatalf("Required environment variable %s is missing", key)
		}
	}

	return &Config{envVars: envVars}
}

func (c *Config) Get(key string) string {
	return c.envVars[key]
}

func (c *Config) Exists(key string) bool {
	_, exists := c.envVars[key]
	return exists
}

func (c *Config) GetInt(key string) int {
	value, _ := strconv.Atoi(c.envVars[key])
	return value
}
