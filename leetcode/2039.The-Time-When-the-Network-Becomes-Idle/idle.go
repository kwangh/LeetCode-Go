package leetcode

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	time := make([]int, n)
	for i := range time {
		time[i] = -1
	}
	time[0] = 0
	qu := []int{0}
	for len(qu) != 0 {
		node := qu[0]
		qu = qu[1:]

		for _, child := range graph[node] {
			if time[child] == -1 {
				time[child] = time[node] + 1
				qu = append(qu, child)
			}
		}
	}

	var res int
	for i := 1; i < n; i++ {
		lastOut := ((time[i]*2 - 1) / patience[i]) * patience[i]
		res = Max(res, time[i]*2+lastOut)
	}
	return res + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
