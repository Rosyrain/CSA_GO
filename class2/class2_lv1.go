package main

import "fmt"

type RN interface {
	reverse(number ReverseNumber, n int)
}

type ReverseNumber struct {
	Number []int
	N      int
}

func (rn ReverseNumber) reverse(number ReverseNumber, n int) {
	MaxLength := number.N
	if n > MaxLength {
		fmt.Println("failure,The value of n exceeds the array length...")
	}
	temp := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if n-i-1 < 0 {
			break
		}
		temp = append(temp, number.Number[n-i-1])
		//fmt.Println(temp)
	}
	for j := 0; j < n; j++ {
		number.Number[j] = temp[j]
	}
	fmt.Println(number.Number)
}

func main() {
	var rn RN
	renversenumber := ReverseNumber{[]int{1, 2, 3, 4, 5}, 5}
	rn = renversenumber
	rn.reverse(renversenumber, 3)
}
