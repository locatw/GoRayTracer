package element

import . "../vector"

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
	return Ray{Origin: origin, Direction: direction.Normalize()}
}
