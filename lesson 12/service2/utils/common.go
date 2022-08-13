package utils

import (
	"math/rand"
	"time"
)

// RandInt 取自给定范围内的随机数
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
