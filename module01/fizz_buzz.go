package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz1(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Print("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}

		if i != n {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func FizzBuzz3(n int) {
	for i := 1; i < n; i++ {
		printFizzBuzz(i)
		fmt.Print(", ")
	}
	printFizzBuzz(n)
	fmt.Println()

}

func printFizzBuzz(n int) {
	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Print("Fizz Buzz")
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("Fizz Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
