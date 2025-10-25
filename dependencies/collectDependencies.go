package dependencies

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"pmv/config_load"
	"pmv/parse"
)

// Сбор зависимостей
func CollectDependencies(cfg *config_load.Config) ([]string, error) {
	if cfg.TestMode {
		return parse.ParseDependenciesFromFile(cfg.FilePath)
	}

	if cfg.RepoURL != "" && (len(cfg.RepoURL) > 8 && cfg.RepoURL[:8] == "https://") {
		fmt.Printf("Получение зависимостей пакета '%s' из %s...\n", cfg.Package, cfg.RepoURL)

		resp, err := http.Get(cfg.RepoURL)
		if err != nil {
			return nil, fmt.Errorf("ошибка HTTP-запроса: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения ответа: %v", err)
		}

		var data PipInfo
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, fmt.Errorf("ошибка разбора JSON: %v", err)
		}

		if len(data.Info.RequiresDist) == 0 {
			fmt.Println("У пакета нет явно указанных зависимостей.")
			return []string{}, nil
		}

		return data.Info.RequiresDist, nil
	}

	fmt.Printf("Эмуляция: ")
	return []string{"requests", "numpy", "pandas", "flask"}, nil
}