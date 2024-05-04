package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	fmt.Println(userNames)

	userNames[0] = "Hey"
	userNames = append(userNames, "John")
	userNames = append(userNames, "Max")
	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.9
	fmt.Println(courseRatings)
	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
