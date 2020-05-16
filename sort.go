package ecs

/********************************************************************
created:    2020-05-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func BinarySearch(data []int, key int) int {
	var left, right = 0, len(data) - 1
	var i = left - 1  // 不变式：data[i]<key
	var j = right + 1 // 不变式：data[j]≥key
	for i+1 != j {    // 可以证明：当j≥i+2时[i,j]的范围是在不断缩小的

		var mid = i + ((j - i) >> 1) // (i+j)/2当数字特别大时有可能溢出
		if data[mid] < key {
			i = mid
		} else
		{
			j = mid
		}
	}

	// 查找不成功时如果需要将其插入的话，则无论j==right+1还是data[j]!=key，实际上j实际指向插入的位置
	if j == right+1 || data[j] != key {
		return ^j
	}

	// 查找成功时，j指向key在data[]中第一次出现的位置
	return j
}

func SortBy(keys []int, values []interface{}) {
	var count = len(keys)
	if keys == nil || count != len(values) {
		return
	}

	for i := 1; i < count; i++ {
		for j := i; j > 0 && keys[j ] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
			values[j], values[j-1] = values[j-1], values[j]
		}
	}
}
