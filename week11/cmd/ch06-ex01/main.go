package main

import "fmt"

func main() {
	subjects := []string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3]
	//subjects[0] = "Java"
	subjectsSlice[0] = "java"
	subjectsSlice = append(subjectsSlice, "Go", "Kotlin", "Database")

	for _, susubject := range subjects {
		fmt.Println(susubject)
	}
	fmt.Println("==========")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
