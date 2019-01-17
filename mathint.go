package mathint

import "math"

func Mod(i, j uint32) uint32 {
	return uint32(math.Mod(float64(i), float64(j)))
}

func Square(i, j int32) int32 {
	return int32(math.Pow(float64(i), float64(j)))
}
