package week02

// https://leetcode-cn.com/problems/degree-of-an-array/
// leetcode 697 数组的度

//map 直接存度，起始位置以及最大长度
func findShortestSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var numsMap = make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if len(numsMap[nums[i]]) == 0 {
			numsMap[nums[i]] = []int{1, i, 1}
		} else {
			numsMap[nums[i]] = []int{
				numsMap[nums[i]][0] + 1,
				numsMap[nums[i]][1],
				i - numsMap[nums[i]][1] + 1,
			}
		}
	}
	var deep, length int
	for _, arr := range numsMap {
		if arr[0] > deep {
			length = arr[2]
			deep = arr[0]
		} else if arr[0] == deep {
			if arr[2] < length {
				length = arr[2]
			}
		}
	}
	return length
}

// map存所有下标
func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var numsMap = make(map[int][]int, len(nums))
	var maxDeep int
	var length int
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = append(numsMap[nums[i]], i)
	}
	for _, dis := range numsMap {
		if len(dis) > maxDeep {
			maxDeep = len(dis)
			length = getDistance(dis)
		} else if len(dis) == maxDeep {
			var tmpDis = getDistance(dis)
			if tmpDis < length {
				length = tmpDis
			}
		}
	}
	return length
}

func getDistance(dis []int) int {
	var left = dis[0]
	var right = dis[len(dis)-1]
	for _, distance := range dis {
		if distance < left {
			left = distance
		}
		if distance > right {
			right = distance
		}
	}
	return right - left + 1
}
