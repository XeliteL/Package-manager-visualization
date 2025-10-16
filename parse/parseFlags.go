package parse

import (
	"flag"
	"fmt"
	"pmv/config_load"
	"pmv/errors"
)

// Парсинг флагов
func ParseFlags() config_load.Config {
	cfgPath := flag.String("config", "", "Путь к JSON-файлу конфигурации")
	flagPackage := flag.String("package", "", "Имя анализируемого пакета")
	flagRepo := flag.String("repo", "", "URL репозитория")
	flagFile := flag.String("file", "", "Путь к тестовому файлу")
	flagTest := flag.Bool("test", false, "Режим тестирования")
	flagVersion := flag.String("version", "", "Версия пакета")
	flagASCII := flag.Bool("ascii", false, "Вывод в ASCII-дереве")
	flagMaxDepth := flag.Int("max-depth", -1, "Глубина анализа (-1 = без ограничений)")
	flag.Parse()

	cfg := config_load.Config{MaxDepth: -1}

	if *cfgPath != "" {
		if err := config_load.LoadConfigFromFile(*cfgPath, &cfg); err != nil {
			errors.ExitWithError(fmt.Errorf("ошибка загрузки файла конфигурации '%s': %w", *cfgPath, err))
		}
	}

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

	return cfg
}