package main

import "fmt"

func main() {
	subjects := []string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[1:3]

	for _, susubject := range subjects {
		fmt.Println(susubject)
	}
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
