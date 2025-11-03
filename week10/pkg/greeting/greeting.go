package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

//func hi() {  //외부에 노출 안되는 함수
func Hi() {
	fmt.Println("Hi~")
}
