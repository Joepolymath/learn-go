package main

import "fmt"

func main() {
	// Maps are synonymous to dictionaries or hash tables
	x := make(map[string]int)
	x["age"] = 53
	x["girls"] = 4
	
	fmt.Println(x)

	// delete key word
	delete(x, "girls")
	fmt.Println(x)
}