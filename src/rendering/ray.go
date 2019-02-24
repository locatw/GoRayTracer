package rendering

import . "../vector"

type Ray struct {
	Origin    Vector
	Direction Vector
}

func CreateRay(origin Vector, direction Vector) Ray {
	return Ray{Origin: origin, Direction: direction.Normalize()}
}
