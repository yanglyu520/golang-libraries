package main

import "fmt"

//defer will run the function until all other functions finish running
func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
