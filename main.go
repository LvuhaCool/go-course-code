package main

import "fmt"

func main() {
	score := 0

	fmt.Println("Get Ready!")

	for i := 0; i < 3; i++ {
		fmt.Println("Вы подлетаете к трубе!")
		fmt.Println("*Подлетаете*")
		fmt.Println("")

		fmt.Println("Вы пролетаете через трубу!")
		fmt.Println("*Пролетаете*")
		fmt.Println("")

		fmt.Println("Вы пролетели через трубу!")
		fmt.Println("*Пролетели*")
		fmt.Println("")

		score++
		fmt.Println("Ваш счет:", score)
		fmt.Println("")
	}
}
