package main

import "github.com/k0kubun/pp"

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
	}

	truthy, ok1 := weather[11]
	falsy, ok2 := weather[30]

	pp.Println(truthy, ok1)
	pp.Println(falsy, ok2)
}
