package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	pp.Println(arr[5], arr[6])
	pp.Println()

	arr = append(
		arr, 8,
	)

	pp.Println(arr)

	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	intSlice := make([]int, 0, 5)

	pp.Println(len(intSlice), cap(intSlice))
}
