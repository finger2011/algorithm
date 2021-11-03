package week05onclass

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// leetcode 215. 数组中的第K个最大元素

//小根堆，维护k个值的小根堆
func findKthLargestByHeap(nums []int, k int) int {
	var h = heap{heapList: []int{0}}
	for i := 0; i < len(nums); i++ {
		h.push(nums[i])
		if i >= k {
			h.pop()
		}
	}
	return h.pop()
}

var kthNum []int

//快排减半
func findKthLargest(nums []int, k int) int {
	kthNum = nums
	return quickSort1(0, len(nums)-1, len(nums)-k)
}

func quickSort1(left, right, target int) int {
	if left >= right {
		return kthNum[left]
	}
	var mid = partition1(left, right)
	if target <= mid {
		return quickSort1(left, mid, target)
	}
	return quickSort1(mid+1, right, target)
}

func partition1(left, right int) int {
	var pVal = kthNum[left]
	for left <= right {
		for left < len(kthNum) && kthNum[left] < pVal {
			left++
		}
		for right < len(kthNum) && kthNum[right] > pVal {
			right--
		}
		if left == right {
			break
		}
		if left < right {
			kthNum[left], kthNum[right] = kthNum[right], kthNum[left]
			left++
			right--
		}
	}
	return right
}
