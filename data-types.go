//package declaration
package main

import "fmt"

// import other packages

// package level variables and types
var globmsg string = "I am global"

// package init function

// main function
func main() {
	// Multiple declaration methods
	/* method 1
	var x int
	var y int
	var msg string
	*/
	/* Method 2
	var x,y int
	var msg string
	*/

	/* Method 3 */
	var (
		x, y int
		msg  string
	)

	x = 10
	y = 20
	msg = "result = "

	/* Method 4 */
	var (
		x1   int    = 10
		y1   int    = 30
		msg1 string = "res = "
	)

	/* Method 5 */
	var (
		x2, y2        = 10, 40
		msg2   string = "res = "
	)

	x3, y3, msg3 := 10, 50, "result = "

	fmt.Println(x, y, msg, x+y)
	fmt.Println(x1, y1, msg1, x1+y1)
	fmt.Println(x2, y2, msg2, x2+y2)
	fmt.Println(x3, y3, msg3, x3+y3)

	var (
		x4, y4 int
		msg4   string
	)
	fmt.Println(x4, y4, msg4)

	// formatted strings
	fmt.Printf("x = %d, y = %d and %s %d\n", x, y, msg, x+y)
	// to print type
	fmt.Printf("x = %T, y = %T and %T %T\n", x, y, msg, x+y)

	// type conversion
	var noInt int = 100
	var noFloat float64 = 100.99
	fmt.Println(noInt, noFloat)
	noFloat = float64(noInt)
	fmt.Println(noInt, noFloat)

	//indefinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println(x)

	//switch construct
	/*
		num := 21
		switch num % 2 {
		case 0:
			fmt.Println("num is even")
		case 1:
			fmt.Println("num is odd")
		}
	*/

	switch num := 21; num % 2 {
	case 0:
		fmt.Println("num is even")
	case 1:
		fmt.Println("num is odd")
	}

	/*
		comment by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 18
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	//using fallthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("n is 0")
		fallthrough
	case 1:
		fmt.Println("n is <= 1")
		fallthrough
	case 2:
		fmt.Println("n is <= 2")
		fallthrough
	case 3:
		fmt.Println("n is <= 3")
		fallthrough
	case 4:
		fmt.Println("n is <= 4")
		fallthrough
	case 5:
		fmt.Println("n is <= 5")
		fallthrough
	case 6:
		fmt.Println("n is <= 6")
	case 7:
		fmt.Println("n is <=7")
		fallthrough
	case 8:
		fmt.Println("n is <= 8")
		fallthrough
	case 9:
		fmt.Println("n is <= 9")
	}

}
