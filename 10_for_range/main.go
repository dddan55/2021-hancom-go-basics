package main

import "fmt"

// for, range
func add(numbers ...int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count
}
func main() {
	result := add(1, 2, 3, 4, 5)
	fmt.Println(result)
}

/* 2번
func add(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}
func main() {
	add(10, 20, 30, 40, 50)
}
*/
