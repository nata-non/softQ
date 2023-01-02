package main

import "fmt"

func FibInt(a int) int {
	if a < 2 {
		return a
	}
	return FibInt(a-1) + FibInt(a-2)
}

func main() {
	fmt.Println("FibInt:", FibInt(30))
}
