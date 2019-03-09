package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/locatw/go-ray-tracer/element"
	. "github.com/locatw/go-ray-tracer/image"
	"github.com/locatw/go-ray-tracer/image/pnm"
	mathex "github.com/locatw/go-ray-tracer/math"
	. "github.com/locatw/go-ray-tracer/rendering"
	. "github.com/locatw/go-ray-tracer/vector"
)

func main() {
	width := 640
	height := 640

	ior := new(float64)
	*ior = 1.5168

	camera :=
		CreateCamera(
			Vector{X: 50.0, Y: 52.0, Z: 295.6},
			Vector{X: 0.0, Y: -0.042612, Z: -1.0},
			CreateAxisVector(YAxis),
			Resolution{Width: width, Height: height},
			mathex.ToRadian(30.0))
	scene :=
		Scene{
			Camera: camera,
			Shapes: []Shape{
				&Sphere{
					Center: Vector{X: 27.0, Y: 16.5, Z: 47.0},
					Radius: 16.5,
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           CreateDefaultColor(Black),
						Specular:          Color{R: 0.999, G: 0.999, B: 0.999},
						IndexOfRefraction: nil,
					},
				},
				&Sphere{
					Center: Vector{X: 73.0, Y: 16.5, Z: 78.0},
					Radius: 16.5,
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           CreateDefaultColor(Black),
						Specular:          Color{R: 0.999, G: 0.999, B: 0.999},
						IndexOfRefraction: ior,
					},
				},
				&Sphere{
					Center: Vector{X: 50.0, Y: 681.6 - 0.27, Z: 81.6},
					Radius: 600.0,
					Material: Material{
						Emission:          MultiplyScalar(1000000.0, CreateDefaultColor(White)),
						Diffuse:           Color{R: 0.75, G: 0.75, B: 0.75},
						Specular:          CreateDefaultColor(Black),
						IndexOfRefraction: nil,
					},
				},
				// top
				&Plane{
					Center: Vector{X: 0.0, Y: 81.6, Z: 0.0},
					Normal: Multiply(-1.0, CreateAxisVector(YAxis)),
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           Color{R: 0.75, G: 0.75, B: 0.75},
						Specular:          CreateDefaultColor(Black),
						IndexOfRefraction: nil,
					},
				},
				// bottom
				&Plane{
					Center: Vector{X: 0.0, Y: 0.0, Z: 0.0},
					Normal: CreateAxisVector(YAxis),
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           Color{R: 0.75, G: 0.75, B: 0.75},
						Specular:          CreateDefaultColor(Black),
						IndexOfRefraction: nil,
					},
				},
				// left
				&Plane{
					Center: Vector{X: 1.0, Y: 0.0, Z: 0.0},
					Normal: CreateAxisVector(XAxis),
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           Color{R: 0.75, G: 0.25, B: 0.25},
						Specular:          CreateDefaultColor(Black),
						IndexOfRefraction: nil,
					},
				},
				// right
				&Plane{
					Center: Vector{X: 99.0, Y: 0.0, Z: 0.0},
					Normal: Multiply(-1.0, CreateAxisVector(XAxis)),
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           Color{R: 0.25, G: 0.25, B: 0.75},
						Specular:          CreateDefaultColor(Black),
						IndexOfRefraction: nil,
					},
				},
				// back
				&Plane{
					Center: Vector{X: 0.0, Y: 0.0, Z: 0.0},
					Normal: CreateAxisVector(ZAxis),
					Material: Material{
						Emission:          CreateDefaultColor(Black),
						Diffuse:           Color{R: 0.75, G: 0.75, B: 0.75},
						Specular:          CreateDefaultColor(Black),
						IndexOfRefraction: nil,
					},
				},
			},
		}
	rayTracer := RayTracer{Scene: scene}

	startTime := time.Now()
	image := rayTracer.Render()
	elapsed := time.Since(startTime)

	fmt.Printf("%0.3f [s]\n", elapsed.Seconds())

	err := pnm.WritePpm("image.ppm", image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
