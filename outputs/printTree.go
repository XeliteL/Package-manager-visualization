package outputs

import (
	"fmt"
	"pmv/dependencies"
)

// Вывод дерева
func PrintTree(node *dependencies.DepNode, prefix string, isLast bool) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	fmt.Println(prefix + connector + node.Name)
	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}
	for i, child := range node.Children {
		PrintTree(child, newPrefix, i == len(node.Children)-1)
	}
}