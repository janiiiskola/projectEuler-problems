package main

import (
	"log"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

func main() {
	defer util.TrackExecutionTime(time.Now())

	n, i, res := 600851475143, 2, 0

	for n > 1 {
		if n%i == 0 {
			res = i
			n = n / i
		}
		i++
	}

	log.Printf("Result: %d", res)
}
