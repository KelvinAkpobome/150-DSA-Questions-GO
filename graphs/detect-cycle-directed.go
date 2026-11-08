package graphs

func HasCycleDirected(n int, edges [][]int) bool {
    adj := make([][]int, n)

    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
    }

    type state int8

    const (
        UNSEEN state = iota
        ACTIVE
        DONE
    )

    states := make([]state, n)

    var dfs func(v int) bool

    dfs = func(v int) bool {
        states[v] = ACTIVE

        for _, u := range adj[v] {
            if states[u] == ACTIVE {
                return true
            }

            if states[u] == UNSEEN && dfs(u) {
                return true
            }
        }

        states[v] = DONE
        return false
    }

    for v := 0; v < n; v++ {
        if states[v] == UNSEEN && dfs(v) {
            return true
        }
    }

    return false
}
