package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	defer func() {
		fmt.Println("5")
	}()
	fmt.Println(secondary())
}

func secondary() int {
	defer func() {
		fmt.Println("3")
	}()
	fmt.Println(2)
	return 4
}
