package utils

import (
	"math/rand"
	"strconv"
)

func RandomNumber6character() string {
	return string(strconv.Itoa(randRange(100000, 999999)))
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
