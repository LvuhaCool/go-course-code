package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[5], arr[6])

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
}
