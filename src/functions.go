package main

import "fmt"

func main() {
	fmt.Println(add(5, 6, 98, 676, 4565))
	fmt.Println(multiple())
	x, y := multiple()
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(factorial(4))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func multiple() (int, int) {
	return 7, 9
}

func factorial(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}