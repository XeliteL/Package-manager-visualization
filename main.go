package main

import (
	"flag"
	"fmt"
	"pmv/config_load"
	"pmv/errors"
	"pmv/validate"
)

func main() {
	// Определяем флаги командной строки
	cfgPath := flag.String("config", "", "Путь к JSON-файлу конфигурации")
	flagPackage := flag.String("package", "", "Имя анализируемого пакета")
	flagRepo := flag.String("repo", "", "URL репозитория")
	flagFile := flag.String("file", "", "Путь к тестовому файлу")
	flagTest := flag.Bool("test", false, "Включить тестовый режим (использует локальный файл)")
	flagVersion := flag.String("version", "", "Версия пакета")
	flagASCII := flag.Bool("ascii", false, "Выводить зависимости в формате ASCII-дерева")
	flagMaxDepth := flag.Int("max-depth", -1, "Максимальная глубина анализа (>=0, -1 = без ограничений)")

	flag.Parse()

	// Создаём базовую конфигурацию со значениями по умолчанию
	cfg := config_load.Config{
		MaxDepth: -1,
	}

	// Если указан JSON-файл — читаем из него
	if *cfgPath != "" {
		if err := config_load.LoadConfigFromFile(*cfgPath, &cfg); err != nil {
			errors.ExitWithError(fmt.Errorf("ошибка загрузки файла конфигурации '%s': %w", *cfgPath, err))
		}
	}

	// Флаги командной строки имеют приоритет над JSON-конфигом
	if *flagPackage != "" {
		cfg.Package = *flagPackage
	}
	if *flagRepo != "" {
		cfg.RepoURL = *flagRepo
	}
	if *flagFile != "" {
		cfg.FilePath = *flagFile
	}
	if *flagTest {
		cfg.TestMode = true
	}
	if *flagVersion != "" {
		cfg.Version = *flagVersion
	}
	if *flagASCII {
		cfg.ASCII = true
	}
	if *flagMaxDepth != -1 {
		cfg.MaxDepth = *flagMaxDepth
	}

	// Проверяем корректность введённых данных
	if err := validate.ValidateConfig(&cfg); err != nil {
		errors.ExitWithError(err)
	}

	// Выводим итоговую конфигурацию в формате key=value
	config_load.PrintConfig(&cfg)
}
