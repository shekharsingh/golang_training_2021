//package declaration
package main

import "fmt"

// main function
func main() {
	// Assignment : Find out the first 20 prime numebrs starting from 5
	prime := 5
	numPrime := 1

	for numPrime <= 21 {
		isPrime := true
		for i := 2; i < prime; i++ {
			if prime%i == 0 {
				isPrime = false
			}
		}
		if isPrime == true {
			fmt.Printf("prime number %d = %d\n", numPrime, prime)
			numPrime += 1
		}
		prime += 1
	}
}
