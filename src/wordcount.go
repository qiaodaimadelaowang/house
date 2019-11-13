package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	fields := strings.Fields(s)
	for _, v := range fields {
		if _, ok := m[v]; ok {
			m[v] += m[v]
		} else {
			m[v] = 1
		}
	}
	return m
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	count := WordCount("sdf a d a d d ads wo owo wo")
	for k, v := range count {
		fmt.Printf("%s,%d\n", k, v)
	}
}
