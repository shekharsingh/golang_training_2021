//package declaration
package main

import "fmt"

// main function
func main() {
	println(add(10, 20))
	println(add1(30, 20))
	println(add2(40, 20))
	println(divide(100, 21))
	q, _ := divide(100, 20)
	println(q)
	// _ is placeholder for variable which we do not want to use.

	// variadic function
	fmt.Println(sum_var(10, 20))
	fmt.Println(sum_var(10, 20, 30))
	fmt.Println(sum_var(10, 20, 30, 40))
	fmt.Println(sum_var(10, 20, 30, 40, 50, 60))

	// Higher order functions
	fn := func() {
		fmt.Println("fn is invoked!")
	}
	fn()
	fmt.Printf("Type of fn is %T\n", fn)

	var addF func(int, int) int
	addF = func(x, y int) int {
		return x + y
	}
	println(addF(100., 200))

	addF = func(x, y int) int {
		return y - x
	}
	println(addF(100., 200))

	// anonymous function
	func() {
		println("anonymous function invoked!")
	}()

	func(x int, y int) {
		println(x / y)
	}(200, 100)

	result := func(x int, y int) int {
		return (x * y)
	}(200, 100)
	println(result)

	logOperation(add2, 34, 54)
	logOperation(subs, 43, 23)

}

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func add2(x, y int) (result int) {
	result = x + y
	return
}

func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

func sum_var(nos ...int) int {
	//fmt.Printf("nos = %#v\n", nos)
	result := 0
	/*
		for idx := 0; idx < len(nos); idx++ {
			result += nos[idx]
		}
	*/
	for _, no := range nos {
		result += no
	}
	return result
}

func logOperation(oper func(int, int) int, x int, y int) {
	fmt.Println("Before Invokation")
	fmt.Println(oper(x, y))
	fmt.Println("After Invokation")
}

func subs(x, y int) int {
	result := x - y
	return result
}
