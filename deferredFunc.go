package main

import "fmt"

func main() {
	/*
		defer func() {
			fmt.Println("Exiting from main!")
		}()
	*/
	defer fmt.Println("Exiting from main!")
	f1()
	fmt.Println("main invoked")
}

func f1() {
	defer fmt.Println("Exiting from f1 - 1")
	defer fmt.Println("Exiting from f1 - 2")
	defer fmt.Println("Exiting from f1 - 3")
	fmt.Println("f1 invoked")
	f2()
}

func f2() {
	defer fmt.Println("Exiting from f2 - 1")
	defer fmt.Println("Exiting from f2 - 2")
	defer fmt.Println("Exiting from f2 - 3")
	fmt.Println("f2 invoked")
}