package main

import (
	"fmt"
)

func main () {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for integer := range ints {
		if integer % 2 == 0 {
			fmt.Println(integer, " is even")
		} else {
			fmt.Println(integer, " is odd")
		}
	}
}