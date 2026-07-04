package main

import (
	"fmt"
)

func main() {
	fmt.Println("До")
	for i := 1; i <= 5; i++ {
		square(i)
	}
	fmt.Println("После")
}

func square(x int) {
	fmt.Println("Квадрат числа:", x*x)
}
