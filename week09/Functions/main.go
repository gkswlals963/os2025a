package main

import "fmt"

<<<<<<< HEAD
func swap(first *int, second *int) {
	temp := 0
=======
<<<<<<< HEAD
func swap(first int, second int) {
	temp := 0
	temp = first
	first = second
	second = temp
	fmt.Println(first, second)
=======
func swap(first *int, second *int) {
	temp := 0
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
	temp = *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
<<<<<<< HEAD
=======
>>>>>>> 0bce5efdf53b901b0878535819a895da8b8a504d
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0

}

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
<<<<<<< HEAD
	swap(&a, &b)
=======
<<<<<<< HEAD
	swap(a, b)
=======
	swap(&a, &b)
>>>>>>> 0bce5efdf53b901b0878535819a895da8b8a504d
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
	fmt.Println(a, b)

}
