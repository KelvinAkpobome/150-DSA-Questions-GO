package graphs

func DFSRecursive(
    graph map[int][]int,
    node int,
    visited map[int]bool,
    order []int,
) []int {
    if visited[node] {
        return order
    }

    visited[node] = true
    order = append(order, node)

    for _, nbr := range graph[node] {
        order = DFSRecursive(graph, nbr, visited, order)
    }

    return order
}

func DFSIterative(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	stack := []int{start}
	result := make([]int, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		visited[node] = true
		result = append(result, node)

		for _, nbr := range graph[node] {
			if !visited[nbr] {
				stack = append(stack, nbr)
			}
		}
	}

	return result
}