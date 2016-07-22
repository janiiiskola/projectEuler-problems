package util

import (
	"log"
	"time"
)

/*
TrackExecutionTime tracks the time since given time.
Usage: defer tools.TrackExecutionTime(time.Now())
*/
func TrackExecutionTime(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
