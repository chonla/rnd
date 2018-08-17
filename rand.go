package rnd

import (
	"math/rand"
	"time"
)

const (
	// ALPHA betic
	ALPHA = "abcdefghijklmnopqrstuvwxyz"
	// NUM eric
	NUM = "0123456789"
	// ALPHANUM eric
	ALPHANUM = "abcdefghijklmnopqrstuvwxyz0123456789"
	// ALPHACASE = alphabetic case sensitive
	ALPHACASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// ALPHANUMCASE = alphanumeric case sensitive
	ALPHANUMCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// HEX adecimal
	HEX = "0123456789abcdef"
	// HEXCASE = hexadecimal case sensitive
	HEXCASE = "0123456789abcdefABCDEF"
)

// OneOf returns one of input
func OneOf(l []interface{}) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return l[r.Intn(len(l))]
}

// Alpha returns a random alphabetic of given length l
func Alpha(l int) string {
	return Of(ALPHA, l)
}

// Alphanum returns a random alphanumeric of given length l
func Alphanum(l int) string {
	return Of(ALPHANUM, l)
}

// Num returns a random numeric of given length l
func Num(l int) string {
	return Of(NUM, l)
}

// Of return a random string in given domain d of given length l
func Of(d string, l int) string {
	dl := len(d)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	o := ""
	for i := 0; i < l; i++ {
		o = o + string(d[r.Intn(dl)])
	}
	return o
}
