package dependencies

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pmv/config_load"
	"pmv/parse"
)

// Сбор зависимостей
func CollectDependencies(cfg *config_load.Config) ([]string, error) {
	if cfg.TestMode {
		return parse.ParseDependenciesFromFile(cfg.FilePath)
	}

	if cfg.RepoURL == "" {
		cfg.RepoURL = fmt.Sprintf("https://pypi.org/pypi/%s/json", cfg.Package)
		fmt.Printf("URL сформирован: %s", cfg.RepoURL)
	}

	fmt.Printf("\nПолучение зависимостей пакета '%s' из %s...\n", cfg.Package, cfg.RepoURL)

	resp, err := http.Get(cfg.RepoURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %v", err)
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