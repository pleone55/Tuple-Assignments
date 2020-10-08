package main

import "fmt"

func main() {
	result := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}(18, 9)

	fmt.Println(result)
}
