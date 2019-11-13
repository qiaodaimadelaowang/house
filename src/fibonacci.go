package main

import "fmt"

func fibonacci() func() int {
	var x = 0
	var y = 0

	return func() int {
		var result = x + y
		if x == 0 && y == 0 {
			y = 1
			return result
		}
		x = y
		y = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
