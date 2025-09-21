package utils

import (
	"math/rand"
	"time"
)

func RandRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func CheckDelay(lastTime time.Time, delay time.Duration) bool {
	return time.Since(lastTime) > delay
}
