package main

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// No need to process nums1 because we are using nums1
	for ; j >= 0; j, k = j-1, k-1 {
		nums1[k] = nums2[j]
	}
}
