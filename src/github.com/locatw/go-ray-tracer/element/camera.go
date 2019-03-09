package element

import . "github.com/locatw/go-ray-tracer/vector"

type Camera struct {
	Origin, Direction, Up Vector
	Fov                   float64
}

func CreateCamera(origin Vector, direction Vector, up Vector, fov float64) Camera {
	correctedDir := Normalize(direction)
	correctedUp := Normalize(Cross(direction, Cross(up, direction)))

	return Camera{Origin: origin, Direction: correctedDir, Up: correctedUp, Fov: fov}
}
