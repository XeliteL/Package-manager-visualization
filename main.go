package main

import (
	"fmt"
	"pmv/dependencies"
	"pmv/errors"
	"pmv/load"
	"pmv/outputs"
	"pmv/parse"
	"pmv/validate"
	"pmv/dots"
	"strings"
)

func main() {
	cfg := parse.ParseFlags()
	if err := validate.ValidateConfig(&cfg); err != nil {
		errors.ExitWithError(err)
	}

	fmt.Println("Режим:", cfg.Mode)
	fmt.Println()

	// Сбор зависимостей
	deps, err := dependencies.CollectDependencies(&cfg)
	if err != nil {
		errors.ExitWithError(fmt.Errorf("ошибка при сборе зависимостей: %w", err))
	}

	// Построение дерева зависимостей
	root := dependencies.BuildDependencyTree(cfg.Package, deps, cfg.MaxDepth)

	// В зависимости от режима — дерево или порядок загрузки
	switch strings.ToLower(cfg.Mode) {
	case "order":
		fmt.Println("Порядок загрузки зависимостей:")
		order := load.GetLoadOrder(root)
		for i, name := range order {
			fmt.Printf("%2d. %s\n", i+1, name)
		}

	case "dot":
		filename := "graph.dot"
		if err := dots.ExportToDOT(root, filename); err != nil {
			errors.ExitWithError(err)
		}
		fmt.Println("Файл визуализации сохранён как:", filename)
		fmt.Println("Можно построить изображение: dot -Tpng graph.dot -o graph.png")

	default:
		fmt.Println("Дерево зависимостей (ASCII):")
		outputs.PrintTree(root, "", true)
	}
}