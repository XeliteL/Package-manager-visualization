package validate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pmv/config_load"
	"strings"
)

// Проверка корректности параметров конфигурации
func ValidateConfig(cfg *config_load.Config) error {
	// Проверка имени пакета
	cfg.Package = strings.TrimSpace(cfg.Package)
	if cfg.Package == "" {
		return errors.New("параметр 'package' не может быть пустым")
	}

	// Проверка корректности URL в случае его наличия
	if cfg.RepoURL != "" {
		if err := ValidateURL(cfg.RepoURL); err != nil {
			return fmt.Errorf("некорректный URL репозитория: %w", err)
		}
	}

	// Проверка наличия файла в случае указания пути до него
	if cfg.FilePath != "" {
		abs, err := filepath.Abs(cfg.FilePath)
		if err != nil {
			return fmt.Errorf("невозможно обработать путь '%s': %w", cfg.FilePath, err)
		}
		info, err := os.Stat(abs)
		if err != nil {
			return fmt.Errorf("файл '%s' не найден или недоступен: %w", abs, err)
		}
		if info.IsDir() {
			return fmt.Errorf("'%s' — это каталог, ожидался файл", abs)
		}
		cfg.FilePath = abs
	}

	// Проверка на наличие файла при включённом тестовом режиме
	if cfg.TestMode && cfg.FilePath == "" {
		return errors.New("в тестовом режиме должен быть указан 'file_path'")
	}

	// Проверяем глубину анализа
	if cfg.MaxDepth < -1 {
		return errors.New("параметр 'max_depth' должен быть равен -1 (без ограничений), либо >= 0")
	}

	return nil
}
