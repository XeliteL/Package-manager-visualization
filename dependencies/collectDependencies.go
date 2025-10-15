package dependencies

import (
	"fmt"
	"pmv/config_load"
)

func CollectDependencies(cfg *config_load.Config) ([]string, error) {
	// Чтение из файла при включённом тестовом режиме
	if cfg.TestMode {
		return ParseDependenciesFromFile(cfg.FilePath)
	}

	// Эмуляция получения зависимостей
	fmt.Printf("Получение зависимостей пакета '%s' из репозитория %s...\n", cfg.Package, cfg.RepoURL)

	// Имитация сетевого запроса и анализа метаданных
	simulated := []string{
		"requests>=2.32.0",
		"numpy>=1.26.0",
		"pandas==2.2.2",
	}
	return simulated, nil
}
