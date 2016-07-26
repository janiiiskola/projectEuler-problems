package main

import (
	"log"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

const (
	min = 2
	max = 20
)

func main() {
	defer util.TrackExecutionTime(time.Now())

	var a, i int64 = 1, min
	for ; i <= max; i++ {
		// Formula for lcm with more than 2 integers:
		// lcm(a, b, c) = lcm(a, lcm(b, lcm(c, d)))
		a = lcm(a, i)
	}
	log.Printf("Result: %d", a)
}

func lcm(a, i int64) int64 {
	return a * i / gcd(a, i)
}

func gcd(a, i int64) int64 {
	// Ref: Euclidean Algorithm
	c := a % i

	if c == 0 {
		return i
	}

	return gcd(i, c)
}
