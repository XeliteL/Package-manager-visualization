package outputs

import (
	"fmt"
	"pmv/config_load"
)

// Вывод конфигурации
func PrintConfig(cfg *config_load.Config) {
	fmt.Printf("package=%s\n", cfg.Package)
	fmt.Printf("repo_url=%s\n", cfg.RepoURL)
	fmt.Printf("file_path=%s\n", cfg.FilePath)
	fmt.Printf("test_mode=%v\n", cfg.TestMode)
	fmt.Printf("version=%s\n", cfg.Version)
	fmt.Printf("ascii=%v\n", cfg.ASCII)
	fmt.Printf("max_depth=%d\n", cfg.MaxDepth)
}