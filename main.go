package main

import "fmt"

func main() {
	score := 177

	if score < 6 || score > 16 {
		fmt.Println("Ты попал в петушиную зону!")
	} else if score > 8 && score < 15 {
		fmt.Println("Ты попал в яблочко!")
	}

	if score >= 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно еще многому научиться!")
	}

	if score == 7 {
		fmt.Println("Ты проиграл!")
	} else if score != 7 {
		fmt.Println("Ты выиграл!")
	}
}
