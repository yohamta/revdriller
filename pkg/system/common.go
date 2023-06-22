package system

import "math/rand"

func randomRotation() float64 {
	return float64(rand.Intn(4) * 90)
}
