package element

import (
	"math"
	"math/rand"

	. "github.com/locatw/go-ray-tracer/image"
	. "github.com/locatw/go-ray-tracer/vector"
)

type Camera struct {
	Origin, Direction, Up Vector
	Fov                   float64
}

func CreateCamera(origin Vector, direction Vector, up Vector, fov float64) Camera {
	correctedDir := Normalize(direction)
	correctedUp := Normalize(Cross(direction, Cross(up, direction)))

	return Camera{
		Origin:    origin,
		Direction: correctedDir,
		Up:        correctedUp,
		Fov:       fov,
	}
}

func (screen *Screen) CreatePixelRays(camera *Camera, x int, y int, samplingCount int) []Ray {
	pixel := screen.createPixel(x, y)

	rays := make([]Ray, samplingCount)
	for i := 0; i < samplingCount; i++ {
		x := rand.Float64() - 0.5
		y := rand.Float64() - 0.5

		subPixelPos := pixel.calculateSubPixelPosition(screen, x, y)
		rays[i] = CreateRay(camera.Origin, Subtract(subPixelPos, camera.Origin))
	}

	return rays
}

type Screen struct {
	Center     Vector
	XAxis      Vector
	YAxis      Vector
	Resolution Resolution
	Width      float64
	Height     float64
}

type pixel struct {
	Center Vector
	Width  float64
	Height float64
}

func (camera *Camera) CreateScreen(resolution Resolution) Screen {
	aspect := resolution.Aspect()

	center := Add(camera.Origin, camera.Direction)
	xAxis := Normalize(Cross(camera.Direction, camera.Up))
	yAxis := Multiply(-1.0, camera.Up)
	height := 2.0 * math.Tan(camera.Fov/2.0)
	width := height * aspect

	return Screen{
		Center:     center,
		XAxis:      xAxis,
		YAxis:      yAxis,
		Resolution: resolution,
		Width:      width,
		Height:     height,
	}
}

func (screen *Screen) createPixel(x int, y int) pixel {
	leftTopPixel := screen.createLeftTopPixel()

	center := Add(leftTopPixel.Center, Multiply(float64(x)*leftTopPixel.Width, screen.XAxis))
	center = Add(center, Multiply(float64(y)*leftTopPixel.Height, screen.YAxis))

	return pixel{Center: center, Width: leftTopPixel.Width, Height: leftTopPixel.Height}
}

func (screen *Screen) createLeftTopPixel() pixel {
	pixelWidth := screen.Width / float64(screen.Resolution.Width)
	pixelHeight := screen.Height / float64(screen.Resolution.Height)

	offset := Multiply(pixelWidth/2.0, screen.XAxis)
	offset = Add(offset, Multiply(pixelHeight/2.0, screen.YAxis))

	center := Subtract(screen.Center, Multiply(screen.Width/2.0, screen.XAxis))
	center = Subtract(center, Multiply(screen.Height/2.0, screen.YAxis))
	center = Add(center, offset)

	return pixel{Center: center, Width: pixelWidth, Height: pixelHeight}
}

func (pixel *pixel) calculateSubPixelPosition(screen *Screen, x float64, y float64) Vector {
	return AddAll(pixel.Center, Multiply(x*pixel.Width, screen.XAxis), Multiply(y*pixel.Height, screen.YAxis))
}
