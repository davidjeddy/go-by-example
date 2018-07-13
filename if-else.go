// Branching with `if` and `else` in Go is
// straight-forward.

package main

import "fmt"

func main() {

	// Here's a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an `if` statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in all
	// branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// defining multiple variables preceding the conditional of an if statement
	if num, num2 := 9, 10; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 || num2 < 10 {
		fmt.Println(num, "has 1 digit #2")
	} else {
		fmt.Println(num2, "has multiple digits #2")
	}
}

// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.
