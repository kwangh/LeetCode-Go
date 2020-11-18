package leetcode

// FindMedianSortedArrays find median of the two sorted arrays
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := 0.0
	len1, len2, index1, index2 := len(nums1), len(nums2), 0, 0

	merged := []int{}
	for index1 < len1 || index2 < len2 {
		if index1 == len1 && index2 < len2 {
			merged = append(merged, nums2[index2])
			index2++
			continue
		}

		if index2 == len2 && index1 < len1 {
			merged = append(merged, nums1[index1])
			index1++
			continue
		}

		if nums1[index1] < nums2[index2] {
			merged = append(merged, nums1[index1])
			index1++
		} else {
			merged = append(merged, nums2[index2])
			index2++
		}
	}

	if len(merged)%2 == 1 {
		res = float64(merged[len(merged)/2])
	} else {
		i := len(merged)/2 - 1
		res = float64(merged[i]+merged[i+1]) / 2
	}

	return res
}
