package dots

import (
	"fmt"
	"os"
	"pmv/dependencies"
)

// Рекурсивная запись узлов в DOT
func WriteDOTNode(f *os.File, node *dependencies.DepNode) {
	for _, child := range node.Children {
		fmt.Fprintf(f, "    \"%s\" -> \"%s\";\n", node.Name, child.Name)
		WriteDOTNode(f, child)
	}
}