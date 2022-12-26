package main

import "fmt"

func main() {
	// First for loop strategy
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("Moving to the second strategy...")

	// Second for loop strategy
	for j:=0; j<10; j++ {
		fmt.Println(j+1)
	}

	number := 5
	if number % 2 == 0 {
		fmt.Printf("Number %d is even", number)
	} else {
		fmt.Printf("Number %d is odd \n", number)
	}

	// Suppose we want to print out the english equivalent of numbers
	k := 4
	switch k {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown")
	}


	// Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
	for m:=1; m<=100; m++ {
		if m % 3 == 0 {
			fmt.Println(m)
		}
	}

	// Go Arrays
	// To calculate average of array elements
	var nums [5]float32
	nums[0] = 24
	nums[1] = 77
	nums[2] = 25
	nums[3] = 62
	nums[4] = 98

	var total float32 = 0
	for b:=0; b < len(nums); b++ {
		total += nums[b]
	}
	fmt.Println("Printing Array Average")
	fmt.Println(total / float32(len(nums)))

	// Another strategy for creating and declaring an array
	array := [5]float32{
		2,5,7,34,25,
	}
	fmt.Println(array)
}