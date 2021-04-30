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
