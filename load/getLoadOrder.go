package load

import "pmv/dependencies"

func GetLoadOrder(root *dependencies.DepNode) []string {
    var order []string
    visited := make(map[string]bool)

    var dfs func(*dependencies.DepNode)
    dfs = func(node *dependencies.DepNode) {
        if visited[node.Name] {
            return
        }
        visited[node.Name] = true
        for _, child := range node.Children {
            dfs(child)
        }
        order = append(order, node.Name)
    }

    dfs(root)
    return order
}
