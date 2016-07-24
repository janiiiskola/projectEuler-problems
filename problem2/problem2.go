package main

import (
	"log"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

func main() {
	defer util.TrackExecutionTime(time.Now())

	a, sum, i := [2]int{1, 1}, 0, 0

	for i < 4e6 {
		i = a[0] + a[1]
		a[0] = a[1]
		a[1] = i

		if i%2 == 0 {
			sum += i
		}

	}

	log.Printf("Result: %d", sum)
}
