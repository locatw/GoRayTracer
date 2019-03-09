package rendering

import . "github.com/locatw/go-ray-tracer/element"

type Scene struct {
	Camera Camera
	Shapes []Shape
}
