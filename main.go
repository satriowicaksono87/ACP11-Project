package main

import "fmt"

func main() {
	fmt.Println("Hello World3")
	fmt.Println("Fixing Bug")
	fmt.Println("Hello World 1")
	fmt.Println("Hello World 2")
	fmt.Println("Hello World 3")
}

func featureA() {
	fmt.Println("This is Feature A")
	fmt.Println("Pull request feature A")
}
func featureB() {
	a := 10
	b := 5
	fmt.Println(a * b)
}
