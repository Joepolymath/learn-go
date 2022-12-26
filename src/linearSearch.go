package main

import "fmt"

// An implemntation of linear search algorithm.

func linearSearch(array []int, element int) bool  {
	isFound := false
	// Looping through to search
	for i:=0; i <= len(array); i++ {
		if element == array[i]{
			isFound = true
			break
		}
	}
	return isFound
}

func main() {
	fmt.Println("THIS PERFORMS LINEAR SEARCH.")
	array := [5]int{
		4,
		6,
		8,
		12,
		65,
	}

	var result bool = linearSearch(array[:], 6)
	if result {
		fmt.Println("Found 😎")
	} else {
		fmt.Println("NOT FOUND 😞")
	}
}