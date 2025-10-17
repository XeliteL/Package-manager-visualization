package parse

import (
	"flag"
	"pmv/config_load"
)

// Парсинг флагов
func ParseFlags() config_load.Config {
	cfg := config_load.Config{MaxDepth: -1, Mode: "tree"}
	configFile := flag.String("config", "", "Путь к JSON-файлу конфигурации")
	flag.StringVar(&cfg.Package, "package", "", "Имя пакета")
	flag.StringVar(&cfg.RepoURL, "repo", "", "URL репозитория")
	flag.StringVar(&cfg.FilePath, "file", "", "Путь к тестовому файлу зависимостей")
	flag.BoolVar(&cfg.TestMode, "test", false, "Режим тестирования")
	flag.BoolVar(&cfg.ASCII, "ascii", true, "Вывод в ASCII-дереве")
	flag.IntVar(&cfg.MaxDepth, "max-depth", -1, "Глубина анализа (-1 = без ограничений)")
	flag.StringVar(&cfg.Mode, "mode", "tree", "Режим работы: tree, order или dot")
	flag.Parse()

	if *configFile != "" {
		_ = config_load.LoadConfigFromFile(*configFile, &cfg)
	}
	return cfg
}