package simple

import "sort"

//leed code 350题两个数的交集

// 给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
//应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]

func Intersect(nums1 []int, nums2 []int) []int {
	mapData := map[int]int{}
	sort.Ints(nums1)
	result := make([]int, 0, len(nums2))
	for i := 0; i < len(nums1); i++ {
		mapData[nums1[i]] += 1
	}
	for i := 0; i < len(nums2); i++ {
		time, ok := mapData[nums2[i]]
		if ok {
			result = append(result, nums2[i])
			if time > 1 {
				mapData[nums2[i]] -= 1
			} else {
				delete(mapData, nums2[i])
			}
		}
	}
	return result
}
