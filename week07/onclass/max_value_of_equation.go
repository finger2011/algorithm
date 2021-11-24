package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/max-value-of-equation/
// leetcode 1499. 满足不等式的最大值

func findMaxValueOfEquation(points [][]int, k int) int {
	var q []int
	var ans = int(-1e9)
	for i := 0; i < len(points); i++ {
		for len(q) > 0 && points[q[0]][0] < points[i][0]-k {
			q = q[1:]
		}
		if len(q) > 0 {
			ans = internel.Max(ans, points[i][1] + points[i][0] + points[q[0]][1] - points[q[0]][0])
		}
		for len(q) > 0 && points[q[len(q) - 1]][1]- points[q[len(q) - 1]][0] <= points[i][1] - points[i][0] {
			q = q[0:len(q) - 1]
		}
		q = append(q, i)
	}
	return ans
}
