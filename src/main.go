package main

import (
	"fmt"
	"os"

	. "./element"
	"./image/pnm"
	. "./rendering"
	. "./vector"
)

func main() {
	width := 64
	height := 64

	camera :=
		CreateCamera(
			Vector{X: 0.0, Y: 0.0, Z: 10.0},
			Vector{X: 0.0, Y: 0.0, Z: -1.0},
			Vector{X: 0.0, Y: 1.0, Z: 0.0},
			60.0)
	scene :=
		Scene{
			Camera: camera,
			Shapes: []Shape{
				&Sphere{Center: Vector{X: 0.0, Y: 0.0, Z: 0.0}, Radius: 1.0},
			},
		}

	image := Render(scene, width, height)

	err := pnm.WritePpm("image.ppm", image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
