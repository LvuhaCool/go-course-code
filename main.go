package main

import "fmt"

func main() {
	number := 1
	for i := 1; i <= 3; i++ {
		inLoop := 1
		fmt.Println("Iteration:", i, "Number:", number, "InLoop:", inLoop)
		if i < 3 {
			number++
		}
	}
	fmt.Println("Final number:", number, "Final inLoop:", "Not accessible outside the loop")
}
