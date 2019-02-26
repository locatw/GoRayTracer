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
			Multiply(-1.0, CreateAxisVector(ZAxis)),
			CreateAxisVector(YAxis),
			60.0)
	scene :=
		Scene{
			Camera: camera,
			Shapes: []Shape{
				&Sphere{Center: CreateZeroVector(), Radius: 1.0},
			},
		}

	image := Render(scene, width, height)

	err := pnm.WritePpm("image.ppm", image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
