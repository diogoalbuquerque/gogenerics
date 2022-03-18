package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(same(1, 2))
	fmt.Println(same(1.12, 2.11))
	fmt.Println(same(1.99, 1.99))
	fmt.Println(miniAndMaxi([]int{1, 5, 3, 2, 4}))
	fmt.Println(onlyOne([]string{"Banana", "Carrot", "Guava", "Pineapple", "Carrot", "Guava"}))
	fmt.Println(onlyOne([]int{1, 1, 10, 9, 9, 5}))
}

func same[T comparable](val1 T, val2 T) bool {
	if val1 == val2 {
		return true
	}
	return false
}

func miniAndMaxi[T constraints.Ordered](collection []T) (T, T) {

	min := collection[0]
	max := collection[0]
	for _, value := range collection {

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func onlyOne[T comparable](collection []T) []T {
	response := make([]T, 0, len(collection))
	input := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := input[item]; ok {
			continue
		}

		input[item] = struct{}{}
		response = append(response, item)
	}

	return response
}
