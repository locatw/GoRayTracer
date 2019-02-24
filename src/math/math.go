package math

import std_math "math"

func Round32(value float32) float32 {
	return float32(std_math.Round(float64(value)))
}

func Clamp32(value float32, min float32, max float32) float32 {
	value64 := float64(value)
	min64 := float64(min)
	max64 := float64(max)

	return float32(std_math.Min(max64, std_math.Max(min64, value64)))
}
