package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	pnt1, pnt2, pnt3 := m-1, n-1, len(nums1)-1

	if pnt1 == -1 {
		for i := range len(nums1) {
			nums1[i] = nums2[i]
		}
		return
	}

	for pnt2 >= 0 {
		if pnt1 >= 0 && nums1[pnt1] > nums2[pnt2] {
			nums1[pnt3] = nums1[pnt1]
			pnt1--
		} else {
			nums1[pnt3] = nums2[pnt2]
			pnt2--
		}
		pnt3--
	}

}
