package main

import "fmt"

func main() {
	// Working with pointers...
	x := 5
	a := 500
	fmt.Println(x)
	fmt.Println(&x)
	change(&x)
	fmt.Println(x)
	change2(a)
	fmt.Println(a)
}

func change(xPtr *int) {
	fmt.Println(*xPtr)
	*xPtr = 897
}

func change2(a int) {
	fmt.Println(a)
	a = 7652
}