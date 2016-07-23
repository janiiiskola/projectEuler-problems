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

	for i := min; i <= max; i++ {
		for y := min; y <= max; y++ {
			product := strconv.Itoa(i * y)
			reversedProduct := reverseString(product)
			if product == reversedProduct {
				n, err := strconv.Atoi(product)
				if err != nil {
					log.Fatal(err)
				}
				if largest == 0 || n > largest {
					largest = n
				}
			}
		}
	}
	log.Printf("Result: %d", largest)
}

// Creates and palindrome of given byte slice
func reverseString(s string) string {
	var buffer bytes.Buffer
	for i := len(s) - 1; i > -1; i-- {
		buffer.WriteString(string(s[i]))
	}
	return buffer.String()
}
