package main

import "fmt"

func main() {
	result := func(n int) int {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			x, y = y, x+y
		}
		return x
	}(8)

	fmt.Println(result)
}
