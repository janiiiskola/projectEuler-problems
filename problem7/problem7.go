// Source: https://github.com/swook/projecteuler-golang/blob/master/007.go

package main

import (
	"log"
	"math"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

func isPrime(x int64) bool {
	if x == 2 {
		return true
	}
	var i, z int64 = 2, int64(math.Sqrt(float64(x)))
	for ; i <= z; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	defer util.TrackExecutionTime(time.Now())
	var i int64 = 3
	n := 1
	for {
		if isPrime(i) {
			n++
		}
		if n == 10001 {
			log.Printf("Result: %d", i)
			break
		}
		i++
	}
}
