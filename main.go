package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	defer func() {
		fmt.Println("4")
	}()
	secondary()
}

func secondary() {
	defer func() {
		fmt.Println("3")
	}()
	fmt.Println(2)
}
