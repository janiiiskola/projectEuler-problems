package main

import (
	"log"
	"time"

	"github.com/janiiiskola/projectEuler-problems/util"
)

func main() {
	defer util.TrackExecutionTime(time.Now())

	a := [2]int{1, 1}
	sum := 0
	i := 1

	for i < 4000000 {
		i = a[0] + a[1]
		a[0] = a[1]
		a[1] = i
		if i%2 == 0 {
			sum += i
		}
	}

	log.Println("Result: %d", sum)

}
