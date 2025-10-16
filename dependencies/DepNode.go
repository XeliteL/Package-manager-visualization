package dependencies

// Узел дерева зависимостей
type DepNode struct {
	Name     string
	Children []*DepNode
}