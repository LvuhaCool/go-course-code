package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("Hello, World!")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
}
