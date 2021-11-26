package week08onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
// leetcode 1334. 阈值距离内邻居最少的城市

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				f[i][j] = 0
			} else {
				f[i][j] = 1e9
			}
		}
	}
	for i := 0; i < len(edges); i++ {
		x, y, z := edges[i][0], edges[i][1], edges[i][2]
		f[x][y], f[y][x] = z, z
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = internel.Min(f[i][j], f[i][k] + f[k][j])
			}
		}
	}
	var minN, ans int
	minN = 1e9
	for i := 0; i < n; i++ {
		k := 0
		for j := 0; j < n ; j++ {
			if i != j && f[i][j] <= distanceThreshold {
				k++
			}
		}
		if k < minN || (k == minN && i > ans) {
			ans = i
			minN = k
		}
	}
	return ans
}
