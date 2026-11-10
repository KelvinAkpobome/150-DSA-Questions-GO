package graphs

func DFSRecursive(
	graph map[int][]int,
	node int,
	visited map[int]bool,
	order []int,
) []int {
	//discover a node
	if visited[node] {
		return order
	}
	//after discovery, mark as seen
	visited[node] = true
	order = append(order, node)

	//explore neigbours
	for _, nbr := range graph[node] {
		//send nbr for  discovery
		order = DFSRecursive(graph, nbr, visited, order)
	}

	return order
}

func DFSIterative(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)

	//discover first node
	stack := []int{start}

	//after discovery, mark first node as seen
	visited[start] = true
	result := make([]int, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node)
		//explore neigbours
		for _, nbr := range graph[node] {
			if !visited[nbr] {
				//discover nbr
				stack = append(stack, nbr)

				//after discovery, mark as seen
				visited[nbr] = true
			}
		}
	}

	return result
}
