//package declaration
package main

// import other packages
import "fmt"

// package level variables and types
var globmsg string = "I am global"

// package init function

// main function
func main() {
	// code
	fmt.Println("Hello World!")
	// var variable_name variable_type
	// can't have unused variables in code at function level, it will not compile
	// exception - of variable is defined at package level
	var msg1 string = "My first variable"
	fmt.Println(msg1)
	// type inference
	var msg2 = "No type mentioned"
	fmt.Println(msg2)
	// this is applicable for only function not at package level
	// msg2 := "Just variable"
	// fmt.Println(msg2)
}

// other functions
