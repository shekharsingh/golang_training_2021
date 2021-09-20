package main

import (
	"errors"
	"fmt"
)

var divisionByZeroError = errors.New("division by zero")

func main() {
	result, err := divide(100, 7)
	if err == divisionByZeroError {
		fmt.Println("Cannot divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong...", err)
		return
	}
	fmt.Println(result)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, divisionByZeroError
	}
	return x / y, nil
}