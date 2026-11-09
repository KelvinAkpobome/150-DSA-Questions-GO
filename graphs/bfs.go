package graphs

func BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)

	visited[start] = true
	queue := []int{start}

	result := make([]int, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, nbr := range graph[node] {
			if !visited[nbr] {
				visited[nbr] = true
				queue = append(queue, nbr)
			}
		}
	}

	return result
}
