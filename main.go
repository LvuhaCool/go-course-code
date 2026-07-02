package main

import "fmt"

func main() {
	score := 17

	if score > 10 && score <= 15 {
		fmt.Println("Ты красавчик!")
	} else if score > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else {
		fmt.Println("Тебе нужно еще многому научиться :)")
	}

	if score <= 10 {
		fmt.Println("Тебе нужно еще многому научиться :)")
	} else {
		if score > 15 {
			fmt.Println("Ты мега-красавчик!")
		} else {
			fmt.Println("Ты красавчик!")
		}
	}

	// Решение автора

	if score > 10 {
		if score > 15 {
			fmt.Println("Ты мега-красавчик!")
		} else {
			fmt.Println("Ты красавчик!")
		}
	} else {
		fmt.Println("Тебе нужно еще многому научиться :)")
	}

	// Второе решение автора

	if score > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else {
		if score > 10 {
			fmt.Println("Ты красавчик!")
		} else {
			fmt.Println("Тебе нужно еще многому научиться :)")
		}
	}

	if score > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else if score > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно еще многому научиться :)")
	}
}
