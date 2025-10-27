package dots

import (
	"fmt"
	"os"
	"pmv/dependencies"
)

// Рекурсивная запись узлов в DOT
func WriteDOTNode(f *os.File, node *dependencies.DepNode) {
	parent := SanitizeName(node.Name)
	for _, dep := range node.Children {
		child := SanitizeName(dep.Name)
		fmt.Fprintf(f, "\"%s\" -> \"%s\";\n", parent, child)
		WriteDOTNode(f, dep)
	}
}