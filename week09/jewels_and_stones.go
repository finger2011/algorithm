package week09

// https://leetcode-cn.com/problems/jewels-and-stones/
// leetcode 771. 宝石与石头

func numJewelsInStones(jewels string, stones string) int {
	var jMap = make(map[byte]bool, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jMap[jewels[i]] = true
	}
	var ans int
	for i := 0; i < len(stones); i++ {
		if _, ok := jMap[stones[i]]; ok {
			ans++
		}
	}
	return ans
}