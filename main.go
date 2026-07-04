package main

import "fmt"

func main() {
	subscribed := true
	if !subscribed {
		fmt.Println("Подпишись!")
	}
	notSubscribed := true
	if !notSubscribed {
		fmt.Println("Ура! Ты подписан!")
	}
}
