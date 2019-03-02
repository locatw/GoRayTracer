package element

import (
	"math"
	"math/rand"

	mathex "../math"
	. "../vector"
)

type Ray struct {
	Origin    Vector
	Direction Vector
}

type HitInfo struct {
	Object   Shape
	Position Vector
	Normal   Vector
	T        float64
}

func CreateRay(origin Vector, direction Vector) Ray {
	return Ray{Origin: origin, Direction: Normalize(direction)}
}

func CreateDiffuseRay(ray Ray, hitInfo *HitInfo) Ray {
	u := Normalize(Cross(hitInfo.Normal, ray.Direction))
	v := Normalize(Cross(u, hitInfo.Normal))
	n := hitInfo.Normal

	r := math.Sqrt(rand.Float64())
	theta := 2.0 * math.Pi * rand.Float64()
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	z := math.Sqrt(math.Max(0.0, 1.0-x*x-y*y))

	origin := Add(hitInfo.Position, Multiply(10000.0*mathex.Epsilon(), hitInfo.Normal))
	dir := AddAll(Multiply(x, u), Multiply(y, v), Multiply(z, n))

	return Ray{Origin: origin, Direction: dir}
}

func CreateReflectRay(ray Ray, hitInfo *HitInfo) Ray {
	dir := Multiply(2.0*Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), hitInfo.Normal)
	dir = Subtract(dir, Multiply(-1.0, ray.Direction))

	origin := Add(hitInfo.Position, Multiply(10000.0*mathex.Epsilon(), hitInfo.Normal))

	return Ray{Origin: origin, Direction: dir}
}
