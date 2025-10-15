package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func ParseDependenciesFromFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла зависимостей: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	var deps[]string
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		deps = append(deps, line)
	}

	return deps, nil
}