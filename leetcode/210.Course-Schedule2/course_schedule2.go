package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make(map[int][]int)
	d := make([]int, numCourses)
	for _, v := range prerequisites {
		g[v[1]] = append(g[v[1]], v[0])
		d[v[0]]++
	}

	var q, res []int
	for i := 0; i < numCourses; i++ {
		if d[i] == 0 {
			q = append(q, i)
			res = append(res, i)
		}
	}

	edges := len(prerequisites)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range g[cur] {
			edges--
			d[v]--
			if d[v] == 0 {
				q = append(q, v)
				res = append(res, v)
			}
		}
	}

	if edges != 0 {
		return []int{}
	}

	return res
}
