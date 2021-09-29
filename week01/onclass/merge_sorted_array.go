package week01onclass

//https://leetcode-cn.com/problems/merge-sorted-array/
//leetcode 88

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	i := m - 1
	j := n - 1
	for ; i >= 0; i-- {
		for ; j >= 0; j-- {
			if nums2[j] >= nums1[i] {
				nums1[i+j+1] = nums2[j]
			} else {
				break
			}
		}
		nums1[i+j+1] = nums1[i]
	}
	if j >= 0 {
		for ; j >= 0; j-- {
			nums1[j] = nums2[j]
		}
	}
}
