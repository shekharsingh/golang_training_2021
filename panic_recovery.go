package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var divisionByZeroError = errors.New("division by zero")

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Something went wrong! Contact the administrator..")
			debug.PrintStack()
			fmt.Println(r)
		}
		fmt.Println("Exiting the app")
	}()
	fmt.Println(divide(100, 0))
}

func divide(x, y int) int {
	if y == 0 {
		panic(divisionByZeroError)
	}
	return x / y
}
