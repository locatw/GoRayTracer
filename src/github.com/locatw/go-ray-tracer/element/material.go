package element

import . "github.com/locatw/go-ray-tracer/image"

type Material struct {
	Emission          Color
	Diffuse           Color
	Specular          Color
	IndexOfRefraction *float64
}

func CreateDefaultMaterial() Material {
	black := CreateDefaultColor(Black)

	return Material{Emission: black, Diffuse: black, Specular: black, IndexOfRefraction: nil}
}
