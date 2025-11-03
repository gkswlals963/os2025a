package main

import (
	"fmt"
)

func main() {
	//numbers := [3]int{-9, 11, 7}
	//for i := 0; i < 3; i++ {
	//fmt.Println(i, numbers[i])

	numbers := [3]int{-9, 11, 7}
	for i, number := range numbers {
		fmt.Println(i, number)

	}
}
