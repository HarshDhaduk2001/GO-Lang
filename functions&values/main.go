package main

import "fmt"

type transformFn func(int) int

func main() {
	number := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&number, double)
	tripled := transformNumbers(&number, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformerFunction(&number)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers1 := transformNumbers(&number, transformFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(val))
	}

	return transformedNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
