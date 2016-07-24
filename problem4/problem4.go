package main

import (
	"bytes"
	"log"
	"strconv"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

const (
	max = 999
	min = 100
)

func main() {
	defer util.TrackExecutionTime(time.Now())

	largest := 0

	// Loop all possible products of different combinations of 100 and 999
	for i := min; i <= max; i++ {
		for y := min; y <= max; y++ {

			// Convert product to string, create a reverted version and compare to original
			product := strconv.Itoa(i * y)
			reversedProduct := reverseString(product)

			// If reverted and original are equal, check if product is largest number so far and store value if largest
			if product == reversedProduct {
				n, err := strconv.Atoi(product)
				if err != nil {
					log.Fatal(err)
				}

				if largest == 0 || n > largest {
					largest = n
				}
			}
			// End for loop y
		}
		// End for loop i
	}

	log.Printf("Result: %d", largest)
}

// reverseString returns a palindrome of given string
func reverseString(s string) string {

	var buffer bytes.Buffer

	for i := len(s) - 1; i > -1; i-- {
		buffer.WriteString(string(s[i]))
	}

	return buffer.String()
}
