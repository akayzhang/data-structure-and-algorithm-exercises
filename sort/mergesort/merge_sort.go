package main

import (
	"fmt"
)

// 归并排序
func MergeSort(a []int64) []int64 {
	l := len(a)
	if l == 1 || l == 0 {
		return a
	}

	median := l / 2
	left := a[0 : median]
	right := a[median: l]

	leftSorted := MergeSort(left)
	rightSorted := MergeSort(right)
	lLen := len(leftSorted)
	rLen := len(rightSorted)
	newIndex := 0
	lIndex := 0
	rIndex := 0
	newSlice := make([]int64, l)
	for lIndex < lLen && rIndex < rLen {
		if leftSorted[lIndex] < rightSorted[rIndex] {
			newSlice[newIndex] = leftSorted[lIndex]
			lIndex++
		} else {
			newSlice[newIndex] = rightSorted[rIndex]
			rIndex++
		}
		newIndex++
	}

	if lIndex == lLen {
		for _, v := range rightSorted[rIndex:] {
			newSlice[newIndex] = v
			newIndex++
		}
	} else if rIndex == rLen {
		for _, v := range leftSorted[lIndex:] {
			newSlice[newIndex] = v
			newIndex++
		}
	}

	return newSlice
}

func main()  {
	a := MergeSort([]int64{4,6,12,6,1,3})
	for _,v := range a {
		fmt.Println(v)
	}
}
