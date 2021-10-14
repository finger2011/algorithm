package week03onclass

// https://leetcode-cn.com/problems/redundant-connection/
// leetcode 684. 冗余连接
var toConnections map[int][]int
var visit map[int]bool

//无边数组生成后再判断是否有环
func findRedundantConnection2(edges [][]int) []int {
	var num int
	for i := 0; i < len(edges); i++ {
		if edges[i][0] > num {
			num = edges[i][0]
		}
		if edges[i][1] > num {
			num = edges[i][1]
		}
	}
	toConnections = make(map[int][]int, num+1)
	for i := 0; i < len(edges); i++ {
		if _, ok := toConnections[edges[i][0]]; ok {
			toConnections[edges[i][0]] = append(toConnections[edges[i][0]], edges[i][1])
		} else {
			toConnections[edges[i][0]] = []int{edges[i][1]}
		}
		if _, ok := toConnections[edges[i][1]]; ok {
			toConnections[edges[i][1]] = append(toConnections[edges[i][1]], edges[i][0])
		} else {
			toConnections[edges[i][1]] = []int{edges[i][0]}
		}

	}
	visit = make(map[int]bool, num+1)
	var hasCycle, list = dfs2(1, 0)
	if !hasCycle {
		return []int{}
	}
	visit = make(map[int]bool, num+1)
	var isStart bool
	for i := 1; i < len(list); i++ {
		if list[i] == list[len(list) - 1] {
			//需要先找到环的起点
			isStart = true
		}
		if isStart {
			visit[list[i]] = true
		}		
	}
	for i := len(edges) - 1; i >= 0; i-- {
		if val, ok := visit[edges[i][0]]; ok && val {
			if val, ok := visit[edges[i][1]]; ok && val {
				return edges[i]
			}
		}
	}
	return []int{}

}

func dfs2(start, father int) (bool, []int) {
	visit[start] = true
	for _, y := range toConnections[start] {
		if y == father {
			continue
		}
		if visit[y] {
			return true, []int{father, start, y}
		}
		var son, tmp = dfs2(y, start)
		if son {
			return true, append([]int{father}, tmp...)
		}
	}
	return false, nil
}

func findRedundantConnection(edges [][]int) []int {
	toConnections = make(map[int][]int, len(edges))
	for i := 0; i < len(edges); i++ {
		if _, ok := toConnections[edges[i][0]]; ok {
			toConnections[edges[i][0]] = append(toConnections[edges[i][0]], edges[i][1])
		} else {
			toConnections[edges[i][0]] = []int{edges[i][1]}
		}
		if _, ok := toConnections[edges[i][1]]; ok {
			toConnections[edges[i][1]] = append(toConnections[edges[i][1]], edges[i][0])
		} else {
			toConnections[edges[i][1]] = []int{edges[i][0]}
		}
		visit = make(map[int]bool, len(edges))
		if dfs(edges[i][0], 0) {
			return edges[i]
		}
	}
	return []int{}
}

func dfs(start, father int) bool {
	visit[start] = true
	for _, y := range toConnections[start] {
		if y == father {
			continue
		}
		if visit[y] {
			return true
		}
		var son = dfs(y, start)
		if son {
			return true
		}
	}
	return false
}
