package graphs

func BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	head := 0

	//discover first node
	queue := []int{start}

	//after discovery, mark as seen
	visited[start] = true

	result := make([]int, 0)

	for head < len(queue) {
		node := queue[head]
		head++
		result = append(result, node)

		//explore neigbours
		for _, nbr := range graph[node] {
			if !visited[nbr] {
				//discover nbr
				queue = append(queue, nbr)

				//after discovery, mark as seen
				visited[nbr] = true

			}
		}
	}

	return result
}
