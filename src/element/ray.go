package element

import (
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

func CreateReflectRay(ray Ray, hitInfo *HitInfo) Ray {
	dir := Multiply(2.0*Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), hitInfo.Normal)
	dir = Subtract(dir, Multiply(-1.0, ray.Direction))

	origin := Add(hitInfo.Position, Multiply(100.0*mathex.Epsilon(), hitInfo.Normal))

	return Ray{Origin: origin, Direction: dir}
}
