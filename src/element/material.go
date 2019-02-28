package element

import . "../image"

type Material struct {
	Emission Color
	Diffuse  Color
}

func CreateDefaultMaterial() Material {
	black := CreateDefaultColor(Black)

	return Material{Emission: black, Diffuse: black, Specular: black}
}
