package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("면적은", int(length)*float64(width))
	fmt.Println("길이 > 너비?", length > float64(width))
	fmt.Println("형변환", reflect.TypeOf(float64(width)))
	fmt.Println("원본", reflect.TypeOf(width))

}
