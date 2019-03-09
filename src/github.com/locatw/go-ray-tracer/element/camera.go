package element

import (
	"math"
	"math/rand"

	"github.com/locatw/go-ray-tracer/image"
	. "github.com/locatw/go-ray-tracer/vector"
)

type Camera struct {
	Origin, Direction, Up Vector
	Fov                   float64
}

func CreateCamera(origin Vector, direction Vector, up Vector, fov float64) Camera {
	correctedDir := Normalize(direction)
	correctedUp := Normalize(Cross(direction, Cross(up, direction)))

	return Camera{Origin: origin, Direction: correctedDir, Up: correctedUp, Fov: fov}
}

func (camera *Camera) CreatePixelRays(width int, height int, coord image.Coordinate, samplingCount int) []Ray {
	aspect := float64(width) / float64(height)

	screenXAxis := Normalize(Cross(camera.Direction, camera.Up))
	screenYAxis := Multiply(-1.0, camera.Up)
	screenHeight := 2.0 * math.Tan(camera.Fov/2.0)
	screenWidth := screenHeight * aspect

	screenCenter := Add(camera.Origin, camera.Direction)

	pixelWidth := screenWidth / float64(width)
	pixelHeight := screenHeight / float64(height)
	offset := Multiply(pixelWidth/2.0, screenXAxis)
	offset = Add(offset, Multiply(pixelHeight/2.0, screenYAxis))

	leftTopPixelCenter := Subtract(screenCenter, Multiply(screenWidth/2.0, screenXAxis))
	leftTopPixelCenter = Subtract(leftTopPixelCenter, Multiply(screenHeight/2.0, screenYAxis))
	leftTopPixelCenter = Add(leftTopPixelCenter, offset)

	pixelCenter := Add(leftTopPixelCenter, Multiply(float64(coord.X)*pixelWidth, screenXAxis))
	pixelCenter = Add(pixelCenter, Multiply(float64(coord.Y)*pixelHeight, screenYAxis))

	rays := make([]Ray, samplingCount)
	for i := 0; i < samplingCount; i++ {
		x := rand.Float64() - 0.5
		y := rand.Float64() - 0.5

		subPixelPos := AddAll(pixelCenter, Multiply(x*pixelWidth, screenXAxis), Multiply(y*pixelHeight, screenYAxis))
		rays[i] = CreateRay(camera.Origin, Subtract(subPixelPos, camera.Origin))
	}

	return rays
}
