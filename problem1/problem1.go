/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
	"log"
	"time"

	"github.com/janiiiskola/projectEuler-problems/tools"
)

func main() {

	// Track the execution time
	defer tools.TrackExecutionTime(time.Now())

	// Store the sum of values
	var sum float64

	// Loop integers below 1000
	var i float64
	for ; i < 1000; i++ {
		// If i divided by 3 or 5 is whole number, increase the sum
		if isWhole(i/3) == true || isWhole(i/5) == true {
			sum += i
		}
	}

	// Print the final sum
	log.Printf("Result: %v", sum)

}

// isWhole returns true if given float is whole number
func isWhole(nro float64) bool {

	if nro == float64(int64(nro)) {
		return true
	} else {
		return false
	}

}
