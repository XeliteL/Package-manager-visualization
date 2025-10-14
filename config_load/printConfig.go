package config_load

import (
	"fmt"
	"strconv"
)

// Вывод конфигурации (ключ = значение)
func PrintConfig(cfg *Config) {
	parameters := map[string]string{
		"package":   cfg.Package,
		"repo_url":  cfg.RepoURL,
		"file_path": cfg.FilePath,
		"test_mode": strconv.FormatBool(cfg.TestMode),
		"version":   cfg.Version,
		"ascii":     strconv.FormatBool(cfg.ASCII),
		"max_depth": strconv.Itoa(cfg.MaxDepth),
	}

	fmt.Println("Итоговая конфигурация:")
	for key, value := range parameters {
		fmt.Printf("%s=%s\n", key, value)
	}
}
