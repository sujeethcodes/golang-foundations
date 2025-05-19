package utils

import "math/rand"

func GenerateNumber() int {
	randNum := rand.Intn(1000)
	return randNum
}
