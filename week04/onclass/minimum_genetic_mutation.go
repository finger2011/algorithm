package week04onclass

// https://leetcode-cn.com/problems/minimum-genetic-mutation/
// leetcode 433. 最小基因变化

func minMutation(start string, end string, bank []string) int {
	var dep = map[string]int{start: 0}
	var hashBank = make(map[string]int, len(bank))
	for i := 0; i < len(bank); i++ {
		hashBank[bank[i]] = 1
	}
	if _, ok := hashBank[end]; !ok {
		return -1
	}
	var codes = []byte{'A', 'C', 'G', 'T'}
	var queue = []string{start}
	for len(queue) > 0 {
		var str = queue[0]
		queue = queue[1:]
		for i := 0; i < len(str); i++ {
			for j := 0; j < len(codes); j++ {
				if codes[j] != str[i] {
					var newStr = str[0:i]
					newStr += string(codes[j])
					if i < len(str)-1 {
						newStr += str[i+1:]
					}
					if _, ok := hashBank[newStr]; !ok {
						continue
					}
					if _, ok := dep[newStr]; ok {
						continue
					}
					queue = append(queue, newStr)
					dep[newStr] = dep[str] + 1
					if newStr == end {
						return dep[newStr]
					}
				}
			}
		}
	}
	return -1
}
