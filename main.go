package main

import "fmt"

func main() {
	score := 17

	if score == 12 {
		fmt.Println("Дюжина")
	} else if score == 21 {
		fmt.Println("Очко")
	} else if score == 50 {
		fmt.Println("Полтинник")
	} else {
		fmt.Println("Ты не попал ни в какую категорию")
	}
}
