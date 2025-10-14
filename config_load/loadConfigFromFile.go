package config_load

import (
	"encoding/json"
	"os"
)

// Загрузка конфигурации из json-файла
func LoadConfigFromFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return err
	}

	return nil
}
