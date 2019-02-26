package element

import . "../vector"

type Camera struct {
	Origin, Direction, Up Vector
	Fov                   float64
}

func CreateCamera(origin Vector, direction Vector, up Vector, fov float64) Camera {
	corrected_dir := Normalize(direction)
	corrected_up := Normalize(Cross(direction, Cross(up, direction)))

	return Camera{Origin: origin, Direction: corrected_dir, Up: corrected_up, Fov: fov}
}
