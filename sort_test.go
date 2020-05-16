package ecs

import (
	"testing"
)

/********************************************************************
created:    2020-05-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func TestSort_BinarySearch(t *testing.T) {

	var data = []int{1, 3, 5, 7, 9, 11}

	var index7 = BinarySearch(data, 7)
	if index7 != 3 {
		t.Errorf("search 7 failed")
	}

	var index4 = BinarySearch(data, 4)
	if index4 != -3 {
		t.Errorf("search 4 failed")
	}
}

func TestSort_SortBy(t *testing.T) {

	type Item struct {
		value int
	}

	var keys = []int{1, 13, 5, 97, 29, 31}
	var values = []interface{}{Item{value: 1}, Item{13}, Item{5}, Item{97}, Item{29}, Item{31}}
	SortBy(keys, values)

	for i := 1; i < len(keys); i++ {
		if keys[i] < keys[i-1] {
			t.Errorf("keys not sorted")
		}

		if keys[i] != values[i].(Item).value {
			t.Errorf("values not sorted")
		}
	}
}
