package rnd

import (
	"math/rand"
	"time"
)

// Random return one of input
func Random(l []interface{}) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return l[r.Intn(len(l))]
}
