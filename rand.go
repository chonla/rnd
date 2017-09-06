package rnd

import (
	"math/rand"
	"time"
)

// OneOf return one of input
func OneOf(l []interface{}) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return l[r.Intn(len(l))]
}
