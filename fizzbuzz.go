package algorithms

import "fmt"

/*
	Given a number, print out numbers from 1 to N
	if number is divisible by 3, print "Fizz"
	if number is divisible by 5, print "Buzz"
	if number is divisible by both 3 and 5, print "Fizz Buzz"
*/

func fizzBuzz(n int) {
	for i := 1; i < n; i++ {
		printFizzBuzz(i)
		fmt.Print(", ")
	}
	printFizzBuzz(n)
	fmt.Println()
}

func printFizzBuzz(n int) {
	switch {
	case n%3 == 0 && n%5 == 0: // can also be divisible by 15
		fmt.Print("Fizz Buzz")
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}
