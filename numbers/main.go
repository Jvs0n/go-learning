package main

import "fmt"

func main() {
	AddNumbers := makeAdder()

	count, sum := AddNumbers(1, 2, 3)
	fmt.Println("Count:", count, "Sum:", sum)

	more := []int{4, 5}
	count, sum = AddNumbers(more...)
	fmt.Println("Count:", count, "Sum:", sum)
}

func makeAdder() func(nums ...int) (int, int) {
	count, sum := 0, 0 // variables captured by closure

	return func(nums ...int) (int, int) {
		for _, n := range nums {
			count++
			sum += n
		}
		return count, sum
	}
}

//Day 2: Variadic Functions, Anonymous Functions, Closures, Multiple return values, name return values, call by value
