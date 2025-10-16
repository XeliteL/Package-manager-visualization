package main

import (
	"fmt"
	"pmv/dependencies"
	"pmv/errors"
	"pmv/outputs"
	"pmv/parse"
	"pmv/validate"
)

func main() {
	cfg := parse.ParseFlags()
	if err := validate.ValidateConfig(&cfg); err != nil {
		errors.ExitWithError(err)
	}

	fmt.Println("Итоговая конфигурация:")
	outputs.PrintConfig(&cfg)
	fmt.Println()

	fmt.Println("Построение дерева зависимостей")
	deps, err := dependencies.CollectDependencies(&cfg)
	if err != nil {
		errors.ExitWithError(fmt.Errorf("ошибка при сборе зависимостей: %w", err))
	}

	root := dependencies.BuildDependencyTree(cfg.Package, deps, cfg.MaxDepth)

	if cfg.ASCII {
		fmt.Println("Дерево зависимостей:")
		outputs.PrintTree(root, "", true)
	} else {
		fmt.Println("Список зависимостей:")
		for _, d := range deps {
			fmt.Println(" -", d)
		}
	}
}