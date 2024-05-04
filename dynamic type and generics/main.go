package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
	result1 := add(1, 2.8)
	fmt.Println(result1)
	result2 := add("1", "2.8")
	fmt.Println(result2)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
