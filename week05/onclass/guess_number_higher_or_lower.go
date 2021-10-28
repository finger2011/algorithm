package week05onclass

// https://leetcode-cn.com/problems/guess-number-higher-or-lower/
// leetcode 374. 猜数字大小
var target int

func guessNumber(n int) int {
    var left, right = 1, n
    for left <= right{
        var mid = (left + right) / 2
        if 0 == guess(mid) {
            return mid
        } else if -1 == guess(mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}

func guess(a int) int {
	if a == target {
		return 0
	} else if a > target {
		return -1
	}
	return 1
}