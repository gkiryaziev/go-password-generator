package rnd

import (
	"math/rand"
)

// get random
func New(seed int64) {
	rand.Seed(seed)
}

// get random between min and max value
func Next(min, max int) int {
	return min + rand.Intn(max-min)
}
