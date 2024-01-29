package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Interface interface {
	Len() int
	Less(i, j int) bool //i,j are indices of sequence elements
	Swap(i, j int)
}

type IntSlice []int

func (arr IntSlice) Len() int {
	return len(arr)
}

func (arr IntSlice) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr IntSlice) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	// 生成100个随机数，append进切片
	var numbers []int
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(1000))
	}

	// 使用sort.Sort和自定义的IntSlice进行排序
	sort.Sort(IntSlice(numbers))

	// 输出排序之后的切片
	fmt.Println("Sorted Numbers:", numbers)
}
