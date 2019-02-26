package vector

import (
	"math/rand"
	"testing"
)

func TestVectorLength(t *testing.T) {
	v := Vector{X: 1.0, Y: 4.0, Z: 8.0}
	expected := 9.0

	result := v.Length()
	if result != expected {
		t.Errorf("%v.Length() must return %f, actual is %f", v, expected, result)
	}
}

func TestVectorAdd(t *testing.T) {
	v1 := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	v2 := Vector{X: 10.0, Y: 20.0, Z: 30.0}
	expected := Vector{X: 11.0, Y: 22.0, Z: 33.0}

	result := Add(v1, v2)
	if result != expected {
		t.Errorf("Add(%v, %v) must return %v, actual is %v", v1, v2, expected, result)
	}
}

func TestVectorAddAll(t *testing.T) {
	v1 := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	v2 := Vector{X: 10.0, Y: 20.0, Z: 30.0}
	v3 := Vector{X: 100.0, Y: 200.0, Z: 300.0}
	expected := Vector{X: 111.0, Y: 222.0, Z: 333.0}

	result := AddAll(v1, v2, v3)
	if result != expected {
		t.Errorf("Add(%v, %v) must return %v, actual is %v", v1, v2, expected, result)
	}
}

func TestVectorSubtract(t *testing.T) {
	v1 := Vector{X: 10.0, Y: 20.0, Z: 30.0}
	v2 := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	expected := Vector{X: 9.0, Y: 18.0, Z: 27.0}

	result := Subtract(v1, v2)
	if result != expected {
		t.Errorf("Subtract(%v, %v) must return %v, actual is %v", v1, v2, expected, result)
	}
}

func TestVectorNormalize(t *testing.T) {
	v := Vector{X: 1.0, Y: 4.0, Z: 8.0}
	v_len := v.Length()

	result := Normalize(v)
	if result.Length() != 1.0 {
		t.Errorf("Normalize(%v) must return a Vector which length is 1.0, actual length is %f", v, result.Length())
	}

	if result.X != v.X/v_len {
		t.Errorf("X value of %v must %f, actual is %f", result, v.X/v_len, result.X)
	}
	if result.Y != v.Y/v_len {
		t.Errorf("Y value of %v must %f, actual is %f", result, v.Y/v_len, result.Y)
	}
	if result.Z != v.Z/v_len {
		t.Errorf("Z value of %v must %f, actual is %f", result, v.Z/v_len, result.Z)
	}
}

func TestVectorDot(t *testing.T) {
	v1 := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	v2 := Vector{X: 10.0, Y: 20.0, Z: 30.0}
	expected := 140.0

	result := Dot(v1, v2)
	if result != expected {
		t.Errorf("Dot(%v, %v) must return %f, actual is %f", v1, v2, expected, result)
	}
}

func TestVectorCross(t *testing.T) {
	patterns := []struct {
		v1, v2   Vector
		expected Vector
	}{
		{
			v1:       Vector{X: 1.0, Y: 0.0, Z: 0.0},
			v2:       Vector{X: 0.0, Y: 1.0, Z: 0.0},
			expected: Vector{X: 0.0, Y: 0.0, Z: 1.0},
		},
		{
			v1:       Vector{X: 0.0, Y: 1.0, Z: 0.0},
			v2:       Vector{X: 0.0, Y: 0.0, Z: 1.0},
			expected: Vector{X: 1.0, Y: 0.0, Z: 0.0},
		},
		{
			v1:       Vector{X: 1.0, Y: 0.0, Z: 0.0},
			v2:       Vector{X: 0.0, Y: 0.0, Z: 1.0},
			expected: Vector{X: 0.0, Y: -1.0, Z: 0.0},
		},
	}

	for _, pattern := range patterns {
		v1 := pattern.v1
		v2 := pattern.v2
		expected := pattern.expected

		result := Cross(v1, v2)
		if result != expected {
			t.Errorf("Cross(%v, %v) must return %f, actual is %f", v1, v2, expected, result)
		}
	}
}

func TestVectorMultiply(t *testing.T) {
	a := 2.0
	v := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	expected := Vector{X: 2.0, Y: 4.0, Z: 6.0}

	result := Multiply(a, v)
	if result != expected {
		t.Errorf("Multiply(%f, %v) must return %v, actual is %v", a, v, expected, result)
	}
}

var loopCount = 1000

func createVecs(count int) []Vector {
	vecs := make([]Vector, count)

	for _, vec := range vecs {
		vec.X = rand.Float64()
		vec.Y = rand.Float64()
		vec.Z = rand.Float64()
	}

	return vecs
}

func BenchmarkVectorAddFunc(b *testing.B) {
	vecs := createVecs(b.N)

	result := CreateZeroVector()

	b.ResetTimer()

	for i := 0; i < loopCount; i++ {
		for j := 0; j < len(vecs); j++ {
			result = Add(result, vecs[j])
		}
	}
	_ = result
}

func BenchmarkVectorAddAllFunc(b *testing.B) {
	vecs := createVecs(b.N)

	result := CreateZeroVector()

	b.ResetTimer()

	for i := 0; i < loopCount; i++ {
		result = AddAll(vecs...)
	}
	_ = result
}
