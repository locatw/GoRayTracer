package element

import (
	"math"
	"math/rand"

	. "github.com/locatw/go-ray-tracer/image"
	. "github.com/locatw/go-ray-tracer/vector"
)

type Camera struct {
	Origin, Direction, Up Vector
	Resolution            Resolution
	Fov                   float64
}

func CreateCamera(origin Vector, direction Vector, up Vector, resolution Resolution, fov float64) Camera {
	correctedDir := Normalize(direction)
	correctedUp := Normalize(Cross(direction, Cross(up, direction)))

	return Camera{
		Origin:     origin,
		Direction:  correctedDir,
		Up:         correctedUp,
		Resolution: resolution,
		Fov:        fov,
	}
}

func (camera *Camera) CreatePixelRays(x int, y int, samplingCount int) []Ray {
	aspect := camera.Resolution.Aspect()

	screenXAxis := Normalize(Cross(camera.Direction, camera.Up))
	screenYAxis := Multiply(-1.0, camera.Up)
	screenHeight := 2.0 * math.Tan(camera.Fov/2.0)
	screenWidth := screenHeight * aspect

	screenCenter := Add(camera.Origin, camera.Direction)

	pixelWidth := screenWidth / float64(camera.Resolution.Width)
	pixelHeight := screenHeight / float64(camera.Resolution.Height)
	offset := Multiply(pixelWidth/2.0, screenXAxis)
	offset = Add(offset, Multiply(pixelHeight/2.0, screenYAxis))

	leftTopPixelCenter := Subtract(screenCenter, Multiply(screenWidth/2.0, screenXAxis))
	leftTopPixelCenter = Subtract(leftTopPixelCenter, Multiply(screenHeight/2.0, screenYAxis))
	leftTopPixelCenter = Add(leftTopPixelCenter, offset)

	pixelCenter := Add(leftTopPixelCenter, Multiply(float64(x)*pixelWidth, screenXAxis))
	pixelCenter = Add(pixelCenter, Multiply(float64(y)*pixelHeight, screenYAxis))

	rays := make([]Ray, samplingCount)
	for i := 0; i < samplingCount; i++ {
		x := rand.Float64() - 0.5
		y := rand.Float64() - 0.5

		subPixelPos := AddAll(pixelCenter, Multiply(x*pixelWidth, screenXAxis), Multiply(y*pixelHeight, screenYAxis))
		rays[i] = CreateRay(camera.Origin, Subtract(subPixelPos, camera.Origin))
	}

	return rays
}
