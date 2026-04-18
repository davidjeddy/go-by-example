// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

package main

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int, int) {
	return 3, 7, 10
}

func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	_, e, _ := vals()
	fmt.Println(e)
}
