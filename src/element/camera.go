package element

import . "../vector"

type Camera struct {
	Origin, Direction, Up Vector
	Fov                   float64
}

func CreateCamera(origin Vector, direction Vector, up Vector, fov float64) Camera {
	corrected_up := Cross(direction, Cross(up, direction))
	corrected_up = corrected_up.Normalize()

	return Camera{Origin: origin, Direction: direction.Normalize(), Up: corrected_up, Fov: fov}
}
