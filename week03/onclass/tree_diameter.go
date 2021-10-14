package week03onclass

// https://leetcode-cn.com/problems/tree-diameter/
// leetcode 1245. 树的直径

var to map[int][]int

func treeDiameter(edges [][]int) int {
	to = make(map[int][]int, len(edges))
	for i := 0; i < len(edges); i++ {
		if _, ok := to[edges[i][0]]; ok {
			to[edges[i][0]] = append(to[edges[i][0]], edges[i][1])
		} else {
			to[edges[i][0]] = []int{edges[i][1]}
		}
		if _, ok := to[edges[i][1]]; ok {
			to[edges[i][1]] = append(to[edges[i][1]], edges[i][0])
		} else {
			to[edges[i][1]] = []int{edges[i][0]}
		}
	}
	p, _ := findFarest(0)
	_, length := findFarest(p)
	return length
}

func findFarest(start int) (int, int) {
	var queue = []int{start}
	var deep = make([]int, len(to))
	// 0未走过
	deep[start] = 1
	for len(queue) > 0 {
		var x = queue[0]
		queue = queue[1:]
		for _, y := range to[x] {
			if deep[y] == 0 {
				deep[y] = deep[x] + 1
				queue = append(queue, y)
			}
		}
	}
	var ans = start
	for i := 0; i < len(deep); i++ {
		if deep[i] > deep[ans] {
			ans = i
		}
	}
	return ans, deep[ans] - 1
}
