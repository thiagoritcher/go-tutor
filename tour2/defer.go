package main

import "fmt"

func main() {
	//Only args are evaluated here
	defer fmt.Println("world")

	fmt.Println("hello")
	//defer is evaluated here
}

