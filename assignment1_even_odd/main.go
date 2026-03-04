package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index := range mySlice {
		if mySlice[index]%2 != 0 {
			fmt.Printf(" %v: Odd", mySlice[index])
		} else {
			fmt.Printf(" %v: Even ", mySlice[index])
		}
	}

}
