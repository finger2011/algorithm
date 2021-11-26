package week08onclass

// https://leetcode-cn.com/problems/min-cost-to-connect-all-points/
// leetcode 1584. 连接所有点的最小费用

var pointFa []int

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	edges := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, []int{i, j, abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
		}
	}
	quickSortPoints(edges, 0, len(edges)-1)
	var ans int
	pointFa = make([]int, len(points))
	for i := 0; i < len(pointFa); i++ {
		pointFa[i] = i
	}
	for i := 0; i < len(edges); i++ {
		x, y, z := edges[i][0], edges[i][1], edges[i][2]
		x, y = findPoint(x), findPoint(y)
		if x != y {
			ans += z
			pointFa[x] = y
		}
	}
	return ans
}

func findPoint(x int) int {
	if pointFa[x] != x {
		pointFa[x] = findPoint(pointFa[x])
	}
	return pointFa[x]
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func quickSortPoints(values [][]int, left, right int) {
	if len(values) <= 1 {
		return
	}
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j][2] >= temp[2] {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i][2] <= temp[2] {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSortPoints(values, left, p-1)
	}
	if right-p > 1 {
		quickSortPoints(values, p+1, right)
	}
}
