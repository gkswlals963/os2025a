package main

import "fmt"

<<<<<<< HEAD
<<<<<<< HEAD
func swap(first *int, second *int) {
	temp := 0
=======
<<<<<<< HEAD
=======
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
func swap(first int, second int) {
	temp := 0
	temp = first
	first = second
	second = temp
	fmt.Println(first, second)
=======
func swap(first *int, second *int) {
	temp := 0
<<<<<<< HEAD
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
=======
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
	temp = *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 0bce5efdf53b901b0878535819a895da8b8a504d
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
=======
>>>>>>> 0bce5efdf53b901b0878535819a895da8b8a504d
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0

}

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
<<<<<<< HEAD
<<<<<<< HEAD
	swap(&a, &b)
=======
<<<<<<< HEAD
=======
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
	swap(a, b)
=======
	swap(&a, &b)
>>>>>>> 0bce5efdf53b901b0878535819a895da8b8a504d
<<<<<<< HEAD
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
=======
>>>>>>> b0bb8e0474618ecba50624d363005e71686f62f0
	fmt.Println(a, b)

}
