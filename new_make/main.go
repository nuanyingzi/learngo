package main

import "fmt"

func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["小西八"] = 10
	fmt.Println(b)
}
