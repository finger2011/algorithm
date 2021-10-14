package week03

// https://leetcode-cn.com/problems/redundant-connection-ii/
// leetcode 685. 冗余连接 II

/*
有根树树+一条有向边，节点数量 = 边数量 = edges长度
因为是有根树+一条有向边，计算所有节点的入度和出度，
如果只有有根树，那么所有节点的入度要么为0（根节点）， 要么为1，加上有向边后有2种情况
1. 所有节点的入度都为1， 该有向边是指向根节点，问题转化为找出根节点，指向根节点的边即为结果
2. 有节点的入度为2，则根节点入度为0， 有向边加在入度为2的节点上
*/
var to map[int][]int
var visit map[int]bool
var inDeg map[int]int

func findRedundantDirectedConnection(edges [][]int) []int {
	to = make(map[int][]int, len(edges)+1)
	inDeg = make(map[int]int, len(edges)+1)
	for i := 0; i < len(edges) + 1; i++ {
		inDeg[i] = 0
	}
	for i := 0; i < len(edges); i++ {
		if _, ok := to[edges[i][0]]; ok {
			to[edges[i][0]] = append(to[edges[i][0]], edges[i][1])
		} else {
			to[edges[i][0]] = []int{edges[i][1]}
		}
		inDeg[edges[i][1]]++
	}
	for i := 1; i < len(inDeg); i++ {
		if inDeg[i] == 2 {
			//有入度为2的节点
			for j := len(edges) - 1; j >= 0; j-- {
				if edges[j][1] == i && checkTree(edges[j]) {
					return edges[j]
				}
			}
		}
	}
	//没有入度为2的节点，则节点一定有环
	visit = make(map[int]bool, len(edges)+1)
	var find = 1
	for i := 2; i < len(to); i++ {
		if len(to[i]) > len(to[find]) {
			find = i
		} 
	}
	var _, list = dfs(find, 0)
	visit = make(map[int]bool, len(edges)+1)
	var isStart bool
	for i := 1; i < len(list); i++ {
		if list[i] == list[len(list)-1] {
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
				if checkTree(edges[i]) {
					return edges[i]
				}
				
			}
		}
	}
	return []int{}
}

func dfs(start, father int) (bool, []int) {
	visit[start] = true
	for _, y := range to[start] {
		// a,b 双向边也算环
		if visit[y] || y == father{
			return true, []int{father, start, y}
		}
		var son, tmp = dfs(y, start)
		if son {
			return true, append([]int{father}, tmp...)
		}
	}
	return false, nil
}

func checkTree(nums []int) bool {
	inDeg[nums[1]]--
	var index int
	for i := 0; i < len(to[nums[0]]); i++ {
		if nums[1] == to[nums[0]][i] {
			index = i
			to[nums[0]][i] = 0
		}
	}

	// var copyDeg = inDeg
	visit = make(map[int]bool, len(inDeg))
	var queue []int
	for i := 1; i < len(inDeg); i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
			break
		}
	}
	var nodes []int
	for len(queue) > 0 {
		var x = queue[0]
		queue = queue[1:]
		nodes = append(nodes, x)
		visit[x] = true
		for _, y := range to[x] {
			if y > 0 && !visit[y] {
				queue = append(queue, y)
			}
			
		}
	}

	to[nums[0]][index] = nums[1]
	inDeg[nums[1]]++
	return len(nodes) == len(inDeg) - 1
}
