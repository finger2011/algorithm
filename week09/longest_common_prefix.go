package week09

// https://leetcode-cn.com/problems/longest-common-prefix
// leetcode 14. 最长公共前缀

func longestCommonPrefix(strs []string) string {
	var ans []byte
	for i := 0; i < len(strs[0]); i++ {
		var same = true
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[j - 1][i] {
				same = false
				break
			}
		}
		if same {
			ans = append(ans, strs[0][i])
		} else {
			break
		}
	}
	return string(ans)
}