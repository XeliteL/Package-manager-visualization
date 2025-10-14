package config_load

type Config struct {
	Package  string `json:"package"`   // Имя пакета
	RepoURL  string `json:"repo_url"`  // URL пакета
	FilePath string `json:"file_path"` // Путь до файла
	TestMode bool   `json:"test_mode"` // Режим тестирования
	Version  string `json:"version"`   // Версия пакета
	ASCII    bool   `json:"ascii"`     // Вывод зависимостей в формате ASCII-дерева.
	MaxDepth int    `json:"max_depth"` // Максимальная глубина анализа зависимостей
}
