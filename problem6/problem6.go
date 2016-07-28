package main

import (
	"log"
	"math"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

const (
	min = 1
	max = 100
	pow = 2
)

func main() {
	defer util.TrackExecutionTime(time.Now())

	var i, p, s float64 = min, 0, 0

	for ; i <= max; i++ {
		s += math.Pow(i, pow)
		p += i
	}

	res := math.Pow(p, pow) - s

	log.Printf("Results: %d", int(res))
}
