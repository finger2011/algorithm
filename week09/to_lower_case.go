package week09

// https://leetcode-cn.com/problems/to-lower-case/
// leetcode 709. 转换成小写字母

func toLowerCase(s string) string {
	var c []byte
	for i := 0; i < len(s); i++ {
		if s[i] < 91 && s[i] > 64 {
			c = append(c, s[i]+32)
		} else {
			c = append(c, s[i])
		}
	}
	return string(c)
}
