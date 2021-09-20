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

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("Unable to divide!")
		}
	}()
	result = divide(x, y)
	return
}

func divide(x, y int) int {
	if y == 0 {
		panic(divisionByZeroError)
	}
	return x / y
}
