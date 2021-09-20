//package declaration
package main

import "fmt"

// main function
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println((sum))

	// Simulate while loop
	nosSum := 1
	for nosSum < 100 {
		nosSum += nosSum
	}
	fmt.Println((nosSum))

	// Assignment : Find out the first 20 prime numebrs starting from 5
	prime := 5
	numPrime := 1

	for numPrime < 21 {
		isPrime := true
		for i := 2; i < prime; i++ {
			if prime%i == 0 {
				isPrime = false
			}
		}
		if isPrime == true {
			fmt.Println(prime)
			numPrime += 1
		}
		prime += 1
	}
}
