package main

import (
	"fmt"
)

type node struct {
	value int64
	point *node
}

// 根据数组生成链表
func NewLinKList(arr []int64) *node {
	var header, current, last *node

	header = &node{}
	current = header
	for _,v := range arr {
		current.value = v
		current.point = &node{}
		last = current
		current = current.point
	}
	last.point = nil
	return header
}

// 打印链表
func PrintLinkList(header *node)  {
	for header != nil {
		fmt.Printf("%s,",header.value)
		header = header.point
	}
	fmt.Println()
}

// 翻转单向链表
func flippingLinkList(header *node) *node {
	current := header
	var last, next ,nextNext *node
	next = current.point


	for current != nil {
		if current.point == nil {
			break
		}

		if last == nil {
			current.point = nil
		}

		nextNext = next.point
		next.point = current
		last = current
		current = next

		if nextNext == nil {
			break;
		}
		next = nextNext
	}
	header = current
	return header
}

func main()  {
	header := NewLinKList([]int64{1,3,5,6})
	PrintLinkList(header)
	header = flippingLinkList(header)
	PrintLinkList(header)

}
