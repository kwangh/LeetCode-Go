package leetcode

func isAcyclic(g map[int][]int, visited, cur map[int]struct{}, node int) bool {
	if _, exist := cur[node]; exist {
		return false
	}

	if _, exist := visited[node]; exist {
		return true
	}

	cur[node] = struct{}{}
	visited[node] = struct{}{}
	for _, v := range g[node] {
		if !isAcyclic(g, visited, cur, v) {
			return false
		}
	}
	delete(cur, node)
	return true
}

func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	degrees := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		degrees[p[0]]++
	}

	var nodes []int
	for i := 0; i < numCourses; i++ {
		if degrees[i] == 0 {
			nodes = append(nodes, i)
		}
	}

	edges := len(prerequisites)
	for len(nodes) != 0 {
		cur := nodes[0]
		nodes = nodes[1:]
		for _, node := range graph[cur] {
			edges--
			degrees[node]--
			if degrees[node] == 0 {
				nodes = append(nodes, node)
			}
		}
	}

	return edges == 0
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	visited := make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		if _, exist := visited[i]; !exist {
			cur := make(map[int]struct{})
			if !isAcyclic(graph, visited, cur, i) {
				return false
			}
		}
	}

	return true
}
