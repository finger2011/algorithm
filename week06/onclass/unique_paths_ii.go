package week06onclasss

// https://leetcode-cn.com/problems/unique-paths-ii/
// leetcode 63. 不同路径 II

//top-down
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var paths = make([][]int, len(obstacleGrid) + 1)
	for i := 0; i < len(paths); i++ {
		paths[i] = make([]int, len(obstacleGrid[0]) + 1)
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				paths[i + 1][j + 1] = 0
			} else if i == 0 && j == 0 {
				paths[i + 1][j + 1] = 1
			} else {
				paths[i + 1][j + 1] = paths[i][j + 1] + paths[i + 1][j] 
			}
		}
	}
	return paths[len(paths) - 1][len(paths[0]) - 1]
}

//bottom-up
var pathGrid [][]int
var path [][]int
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	pathGrid = obstacleGrid
	path = make([][]int, len(obstacleGrid))
	for i := 0; i < len(path); i++ {
		path[i] = make([]int, len(obstacleGrid[0]))
		for j := 0; j < len(path[i]); j++ {
			path[i][j] = -1
		}
	}
	return findPath(0, 0)
}

func findPath(i, j int) int {
	if i > len(pathGrid) - 1 || j > len(pathGrid[0]) - 1 {
		return 0
	}
	if pathGrid[i][j] == 1 {
		return 0
	}
	if i == len(pathGrid) - 1 && j == len(pathGrid[0]) - 1 {
		return 1
	}
	if path[i][j] != -1 {
		return path[i][j]
	}
	ans := findPath(i + 1, j) + findPath(i, j + 1)
	path[i][j] = ans
	return ans
}