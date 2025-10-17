package dots

import (
	"os"
	"pmv/dependencies"
)

// Экспорт графа в формат DOT
func ExportToDOT(root *dependencies.DepNode, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("digraph Dependencies {\n")
	WriteDOTNode(f, root)
	f.WriteString("}\n")
	return nil
}