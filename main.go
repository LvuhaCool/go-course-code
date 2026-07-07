package main

import (
	"fmt"
)

func main() {
	number := 10

	pointer := &number

	fmt.Println(pointer)

	foo(pointer)
}

func foo(n *int) {
	*n += 1
	fmt.Println(*n)
}
