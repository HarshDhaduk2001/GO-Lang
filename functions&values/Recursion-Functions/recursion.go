package main

import "fmt"

func main() {
	fmt.Println(factorial1(4))
	fmt.Println(factorial2(5))
}

// normal function
func factorial1(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

// recursive function
func factorial2(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial2(number-1)
}
