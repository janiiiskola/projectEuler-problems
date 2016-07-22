package util

import (
	"log"
	"time"
)

/*
TrackExecutionTime tracks the time since given time.
Should be used at the beginning of main func:

  defer tools.TrackExecutionTime(time.Now())

*/
func TrackExecutionTime(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
