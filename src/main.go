package main

import (
	"fmt"
	"os"
	"time"

	. "./element"
	. "./image"
	"./image/pnm"
	mathex "./math"
	. "./rendering"
	. "./vector"
)

func main() {
	width := 640
	height := 640

	camera :=
		CreateCamera(
			Vector{X: 50.0, Y: 52.0, Z: 295.6},
			Vector{X: 0.0, Y: -0.042612, Z: -1.0},
			CreateAxisVector(YAxis),
			mathex.ToRadian(30.0))
	scene :=
		Scene{
			Camera: camera,
			Shapes: []Shape{
				&Sphere{
					Center: Vector{X: 27.0, Y: 16.5, Z: 47.0},
					Radius: 16.5,
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: Color{R: 0.999, G: 0.999, B: 0.999},
					},
				},
				&Sphere{
					Center: Vector{X: 73.0, Y: 16.5, Z: 78.0},
					Radius: 16.5,
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: Color{R: 0.999, G: 0.999, B: 0.999},
					},
				},
				&Sphere{
					Center: Vector{X: 50.0, Y: 681.6 - 0.27, Z: 81.6},
					Radius: 600.0,
					Material: Material{
						Emission: MultiplyScalar(1000000.0, CreateDefaultColor(White)),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: CreateDefaultColor(Black),
					},
				},
				// top
				&Plane{
					Center: Vector{X: 0.0, Y: 81.6, Z: 0.0},
					Normal: Multiply(-1.0, CreateAxisVector(YAxis)),
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: CreateDefaultColor(Black),
					},
				},
				// bottom
				&Plane{
					Center: Vector{X: 0.0, Y: 0.0, Z: 0.0},
					Normal: CreateAxisVector(YAxis),
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: CreateDefaultColor(Black),
					},
				},
				// left
				&Plane{
					Center: Vector{X: 1.0, Y: 0.0, Z: 0.0},
					Normal: CreateAxisVector(XAxis),
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: CreateDefaultColor(Black),
					},
				},
				// right
				&Plane{
					Center: Vector{X: 99.0, Y: 0.0, Z: 0.0},
					Normal: Multiply(-1.0, CreateAxisVector(XAxis)),
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: CreateDefaultColor(Black),
					},
				},
				// back
				&Plane{
					Center: Vector{X: 0.0, Y: 0.0, Z: 0.0},
					Normal: CreateAxisVector(ZAxis),
					Material: Material{
						Emission: CreateDefaultColor(Black),
						Diffuse:  Color{R: 0.75, G: 0.75, B: 0.75},
						Specular: CreateDefaultColor(Black),
					},
				},
			},
		}

	start_time := time.Now()
	image := Render(scene, width, height)
	elapsed := time.Since(start_time)

	fmt.Printf("%0.3f [s]\n", elapsed.Seconds())

	err := pnm.WritePpm("image.ppm", image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
