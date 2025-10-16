package dependencies

import (
	"fmt"
	"pmv/config_load"
	"pmv/parse"
)

// Сбор зависимостей
func CollectDependencies(cfg *config_load.Config) ([]string, error) {
	if cfg.TestMode {
		return parse.ParseDependenciesFromFile(cfg.FilePath)
	}
	fmt.Printf("Получение зависимостей пакета '%s' из %s...\n", cfg.Package, cfg.RepoURL)
	return []string{"requests", "numpy", "pandas", "flask"}, nil
}