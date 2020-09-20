package OgameUtil

import (
	"math/rand"
	"time"
)

func RandRange(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}
