package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// LV1 实现一个计算器
func calculator(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

// LV2 生成100个随机数，append进切片，用go自带的函数排序，输出排序之后的切片
func levelTwo() []int {
	var numbers []int
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(1000))
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
	return numbers
}

// LV3 实现 交换排序
func customSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length-1; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			temp := data[min]
			data[min] = data[i]
			data[i] = temp
		}
	}
}

func main() {
	// LV1
	result := calculator(10, 5, "+")
	fmt.Println("LV1 Calculator:", result)

	// LV2
	fmt.Println("LV2 Random Numbers:")
	numbers := levelTwo()

	// LV3 使用自定义排序函数
	customSort(sort.IntSlice(numbers))
	fmt.Println("LV3 Custom Sorted Slice:", numbers)
}
