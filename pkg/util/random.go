package util

import (
	"math/rand"
	"time"
)

// [0,n)
func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// [0,n)
func RandomInt32(n int32) int32 {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Int31n(n)
}

// [0,n)
func RandomInt64(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Int63n(n)
}
