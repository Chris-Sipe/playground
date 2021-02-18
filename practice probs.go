package main

import (
	"fmt"
	"strconv"
)

func notmain() {
	// array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	// sequence := []int{1, 6, -1, 10}
	// if IsValidSubsequence(array, sequence) {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }

	Fizzbuzz()

}

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.

	arrayIndex := 0
	sequenceIndex := 0
	found := false

	for !found {

		if sequenceIndex >= len(sequence) || arrayIndex >= len(array) {
			return false
		}
		// fmt.Printf("comparing %v and %v\n", sequence[sequenceIndex], array[arrayIndex])
		if sequence[sequenceIndex] == array[arrayIndex] {
			// fmt.Println("found match")
			arrayIndex++
			sequenceIndex++
		} else {
			arrayIndex++
		}
		if sequenceIndex >= len(sequence) {
			return true
		}

	}
	return false
}

func Fizzbuzz() {

	for i := 1; i <= 100; i++ {
		// if i%3 == 0 && i%5 == 0 {
		// 	fmt.Printf("%v Fizzbuzz\n", i)
		// } else if i%3 == 0 {
		// 	fmt.Printf("%v Fizz\n", i)
		// } else if i%5 == 0 {
		// 	fmt.Printf("%v Buzz\n", i)
		// } else {
		// 	fmt.Printf("%v\n", i)
		// }

		// output := ""
		// if i%3 == 0 {
		// 	output += "fizz"
		// }
		// if i%5 == 0 {
		// 	output += "buzz"
		// }
		// if output == "" {
		// 	output += strconv.Itoa(i)
		// }

		// fmt.Println(output)

		output := ""
		output += fizzbuzzCheck(i, 3, "fizz")
		output += fizzbuzzCheck(i, 5, "buzz")
		if output == "" {
			output += strconv.Itoa(i)
		}
		fmt.Println(output)

	}

}

func fizzbuzzCheck(i int, n int, str string) string {
	if i%n == 0 {
		return str
	}
	return ""
}
