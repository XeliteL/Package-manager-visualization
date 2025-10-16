package dependencies

// Построение дерева зависимостей
func BuildDependencyTree(rootName string, deps []string, maxDepth int) *DepNode {
	root := &DepNode{ Name: rootName }
	if maxDepth == 0 {
		return root
	}

	for _, dep := range deps {
		child := &DepNode{ Name: dep }
		if maxDepth != 1 {
			child.Children = EmulateSubDeps(dep)
		}
		root.Children = append(root.Children, child)
	}

	return root
}