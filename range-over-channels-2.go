// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan int, 25)

	queue <- 1
	queue <- 2
	queue <- 3
	queue <- 4
	queue <- 5
	queue <- 6
	queue <- 7
	queue <- 8
	queue <- 9
	queue <- 10

	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 10 elements.
	// Shows we can read a channel's contents even after it has been closed.
	for elem := range queue {
		fmt.Println(elem)
	}
}
