package vector

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v *Vector) Length() float64 {
	return math.Sqrt(Dot(*v, *v))
}

func (v Vector) String() string {
	return fmt.Sprintf("Vec(%f, %f, %f)", v.X, v.Y, v.Z)
}

func Add(v1 Vector, v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y, Z: v1.Z + v2.Z}
}

func AddAll(vs ...Vector) Vector {
	result := Vector{X: 0.0, Y: 0.0, Z: 0.0}

	for i := 0; i < len(vs); i++ {
		result.X += vs[i].X
		result.Y += vs[i].Y
		result.Z += vs[i].Z
	}
	return result
}

func Subtract(v1 Vector, v2 Vector) Vector {
	return Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z}
}

func Normalize(v Vector) Vector {
	len := v.Length()
	return Vector{X: v.X / len, Y: v.Y / len, Z: v.Z / len}
}

func Dot(v1 Vector, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Cross(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Multiply(a float64, v Vector) Vector {
	return Vector{X: a * v.X, Y: a * v.Y, Z: a * v.Z}
}
