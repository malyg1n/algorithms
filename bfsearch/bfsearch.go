package bfsearch

type Node struct {
	Children []*Node
	Name     string
}

func Search(node *Node, search string) *Node {
	if node == nil {
		return nil
	}
	var queue []*Node
	visited := make(map[string]bool)
	queue = append(queue, node)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, ok := visited[current.Name]; ok {
			continue
		}
		visited[current.Name] = true
		for _, child := range current.Children {
			if child.Name == search {
				return child
			}
			queue = append(queue, child)
		}
	}

	return nil
}
