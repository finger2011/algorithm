package week05onclass

// https://leetcode-cn.com/problems/sort-an-array/
// leetcode 912. 排序数组

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	var mid = partition(nums,left, right)
	quickSort(nums, left, mid)
	quickSort(nums, mid+1, right)
}

func partition(nums []int,left, right int) int {
	var pVal = nums[left]
	for left <= right {
		for left < len(nums) && nums[left] < pVal {
			left++
		}
		for right < len(nums) && nums[right] > pVal {
			right--
		}
		if left == right {
			break
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return right
}

//归并排序
var numSortArr []int

func sortArray(nums []int) []int {
	numSortArr = nums
	mergeSort(0, len(nums)-1)
	return numSortArr
}
func mergeSort(left, right int) {
	if left >= right {
		return
	}
	var mid = (left + right) / 2
	mergeSort(left, mid)
	mergeSort(mid+1, right)
	merge(left, mid, right)
}

func merge(left, mid, right int) {
	var tmp = make([]int, right-left+1)
	var i, j = left, mid + 1
	for k := 0; k < len(tmp); k++ {
		if j > right || (i <= mid && numSortArr[i] <= numSortArr[j]) {
			tmp[k] = numSortArr[i]
			i++
		} else {
			tmp[k] = numSortArr[j]
			j++
		}
	}
	for k := 0; k < len(tmp); k++ {
		numSortArr[left+k] = tmp[k]
	}
}

// 堆排序
type heap struct {
	heapList []int
}

func (h *heap) pop() int {
	var ans = h.heapList[1]
	h.heapList[1] = h.heapList[len(h.heapList)-1]
	h.heapList = h.heapList[0 : len(h.heapList)-1]
	h.heapifyDown(1)
	return ans
}

func (h *heap) heapifyDown(i int) {
	var ch = 2 * i
	for ch < len(h.heapList) {
		if ch+1 < len(h.heapList) && h.heapList[ch+1] < h.heapList[ch] {
			ch++
		}
		if h.heapList[ch] < h.heapList[i] {
			var tmp = h.heapList[ch]
			h.heapList[ch] = h.heapList[i]
			h.heapList[i] = tmp
			i = ch
			ch = i * 2
		} else {
			break
		}
	}
}

func (h *heap) push(n int) error {
	h.heapList = append(h.heapList, n)
	return h.heapifyUp(len(h.heapList) - 1)
}

func (h *heap) heapifyUp(i int) error {
	for i > 1 {
		if h.heapList[i] < h.heapList[i/2] {
			var tmp = h.heapList[i]
			h.heapList[i] = h.heapList[i/2]
			h.heapList[i/2] = tmp
			i /= 2
		} else {
			break
		}
	}
	return nil
}

func heapSort(nums []int) []int {
	var h = heap{heapList: []int{0}}
	for i := 0; i < len(nums); i++ {
		h.push(nums[i])
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = h.pop()
	}
	return nums
}

//选择排序
func selectionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		var target = i
		for j := i + 1; j < len(nums); j++ {
			if nums[target] > nums[j] {
				target = j
			}
		}
		if target != i {
			var tmp = nums[i]
			nums[i] = nums[target]
			nums[target] = tmp
		}
	}
	return nums
}

//插入排序
func insertionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var ans []int
	for i := 0; i < len(nums); i++ {
		if len(ans) == 0 {
			ans = []int{nums[i]}
		} else {
			var sortable bool
			for j := len(ans) - 1; j >= 0; j-- {
				if nums[i] > ans[j] {
					if j == len(ans)-1 {
						ans = append(ans, nums[i])
					} else {
						ans = append(append(append([]int{}, ans[0:j+1]...), nums[i]), ans[j+1:]...)
					}
					sortable = true
					break
				}
			}
			if !sortable {
				ans = append([]int{nums[i]}, ans...)
			}
		}
	}
	return ans
}

//冒泡排序
func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if nums[j] < nums[j-1] {
				var tmp = nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}
