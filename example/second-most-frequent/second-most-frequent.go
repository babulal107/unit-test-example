package main

import (
	"fmt"
)

/*
   Go program for
   Find the second most frequent character in a given string
*/

func secondMostFrequent(text string) {
	var n int = len(text)
	if n == 0 {
		return
	}
	var first int = 0
	var second int = -1
	var record = make([]int, 256)
	// Count frequency of character element
	for i := 0; i < n; i++ {
		record[text[i]] = record[text[i]]+1
	}

	// Find second most occurring character
	for i := 0; i < n; i++ {
		if record[text[i]] > record[text[first]] {
			second = first
			first = i
		} else if record[text[i]] < record[text[first]] {
			if second == -1 || record[text[second]] < record[text[i]] {
				second = i
			}
		}
	}
	fmt.Println("\n Given Text : ", text)
	if second == -1 {
		fmt.Println("\n Second most frequent element not exist")
	} else {
		fmt.Println("\n Second most frequent element is : ", string(text[second]))
	}
}
func main() {

	// Given Sting text
	var text1 string = "aaaacccff"
	secondMostFrequent(text1)
}
