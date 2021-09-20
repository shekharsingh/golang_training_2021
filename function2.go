// ASSIGNMENT 2
package main

import "fmt"

func main() {
	increment, decrement := getIncrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(decrement()) // => 3
	fmt.Println(decrement()) // => 2
	fmt.Println(decrement()) // => 1
	fmt.Println(decrement()) // => 0
	fmt.Println(decrement()) // => -1

}

func getIncrementer() (func() int, func() int) { //step 1
	var counter int = 0       // step 2
	increment := func() int { //step 3
		counter++ //step 4
		return counter
	}
	decrement := func() int {
		counter--
		return counter
	}
	return increment, decrement //step 5
}