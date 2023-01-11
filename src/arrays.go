package main

import "fmt"

func main(){
	ages := []int{
		5,
		12,
		15,
		18,
		21,
	}
	// Working with slices
	// slice1 := ages[:4]
	slice2 := append(ages, 89, 42)
	anotherArray := make([]int, 2)
	copy(anotherArray, slice2)
	fmt.Printf("The array is %v \n", slice2)
	fmt.Printf("The array is %v \n", anotherArray)
	// fmt.Printf("The array is %v \n", slice3)
}