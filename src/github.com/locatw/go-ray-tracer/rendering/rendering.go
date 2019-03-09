package rendering

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"

	. "github.com/locatw/go-ray-tracer/element"
	"github.com/locatw/go-ray-tracer/image"
	. "github.com/locatw/go-ray-tracer/vector"
)

func distanceAttenuation(ray Ray, hitInfo *HitInfo, color image.Color) image.Color {
	v := Subtract(hitInfo.Position, ray.Origin)
	distance := v.Length()

	return image.DivideScalar(color, 1.0+0.01*math.Pow(distance, 2))
}

func reflectance(ray Ray, normal Vector, n1 float64, n2 float64) float64 {
	// schilick's approximation
	cosTheta := Dot(Multiply(-1.0, ray.Direction), normal)
	r := math.Pow((n1-n2)/(n1+n2), 2.0)

	return r + (1.0-r)*math.Pow(1.0-cosTheta, 5)
}

func traceRay(scene Scene, ray Ray, depth int) image.Color {
	if depth <= 0 {
		return image.CreateDefaultColor(image.Black)
	}

	hitInfo := scene.LookForIntersectedObject(ray)

	if hitInfo == nil {
		return image.CreateDefaultColor(image.Black)
	}

	material := hitInfo.Object.GetMaterial()

	emissionColor := image.CreateDefaultColor(image.Black)
	if !material.Emission.NearlyEqual(emissionColor) {
		emissionColor =
			image.MultiplyScalar(Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), material.Emission)
		emissionColor = distanceAttenuation(ray, hitInfo, emissionColor)
	}

	diffuseColor := image.CreateDefaultColor(image.Black)
	if !material.Diffuse.NearlyEqual(diffuseColor) {
		diffuseRay := CreateDiffuseRay(ray, hitInfo)
		diffuseColor = traceRay(scene, diffuseRay, depth-1)
		diffuseColor = image.MultiplyColor(material.Diffuse, diffuseColor)
		diffuseColor = distanceAttenuation(ray, hitInfo, diffuseColor)
	}

	refractionColor := image.CreateDefaultColor(image.Black)
	refracted := true
	if material.IndexOfRefraction != nil {
		refractRay, isTotalReflection := CreateRefractRay(ray, hitInfo)
		if !isTotalReflection {
			inObject := Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal) < 0.0

			normal := hitInfo.Normal
			if inObject {
				normal = Multiply(-1.0, hitInfo.Normal)
			}

			kr := reflectance(ray, normal, 1.0, *material.IndexOfRefraction)

			if kr < rand.Float64() {
				refractionColor = traceRay(scene, refractRay, depth-1)
				refractionColor = distanceAttenuation(ray, hitInfo, refractionColor)
			} else {
				refracted = false
			}
		} else {
			refracted = false
		}
	} else {
		refracted = false
	}

	specularColor := image.CreateDefaultColor(image.Black)
	if !refracted && !material.Specular.NearlyEqual(specularColor) {
		reflectRay := CreateReflectRay(ray, hitInfo)
		specularColor = traceRay(scene, reflectRay, depth-1)
		specularColor = image.MultiplyColor(material.Specular, specularColor)
		specularColor = distanceAttenuation(ray, hitInfo, specularColor)
	}

	return image.AddColorAll(emissionColor, diffuseColor, specularColor, refractionColor)
}

func createPixelRays(camera Camera, width int, height int, coord image.Coordinate, samplingCount int) []Ray {
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

func renderPixel(scene Scene, width int, height int, pixel *image.Pixel) {
	samplingCount := 1000

	pixelColor := image.CreateDefaultColor(image.Black)
	for _, ray := range createPixelRays(scene.Camera, width, height, pixel.Coordinate, samplingCount) {
		color := traceRay(scene, ray, 10)

		pixelColor = image.AddColor(pixelColor, color)
	}

	pixel.Color = toneMap(image.DivideScalar(pixelColor, float64(samplingCount)))
}

func toneMap(color image.Color) image.Color {
	e := 1.0 / 2.2
	return image.Color{
		R: float32(math.Pow(float64(color.R), e)),
		G: float32(math.Pow(float64(color.G), e)),
		B: float32(math.Pow(float64(color.B), e)),
	}
}

func renderPixelRoutine(pixelCh <-chan *image.Pixel, resultCh chan<- *image.Pixel, scene Scene, width int, height int) {
	for {
		pixel, ok := <-pixelCh
		if !ok {
			break
		}

		renderPixel(scene, width, height, pixel)

		resultCh <- pixel
	}
}

func Render(scene Scene, width int, height int) image.Image {
	rand.Seed(time.Now().UnixNano())

	img := image.CreateImage(width, height)

	pixelCh := make(chan *image.Pixel, width*height)
	resultCh := make(chan *image.Pixel, width*height)

	for i := 0; i < runtime.NumCPU(); i++ {
		go renderPixelRoutine(pixelCh, resultCh, scene, width, height)
	}

	for i := 0; i < len(img.Pixels); i++ {
		pixelCh <- &img.Pixels[i]
	}

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n\n", runtime.NumGoroutine())

	progressPrinter := ProgressPrinter{TotalCount: width * height, Interval: 5 * width, Count: 0}
	finishedPixelCount := 0
	for {
		if finishedPixelCount == width*height {
			break
		}

		<-resultCh

		finishedPixelCount++

		progressPrinter.Print()
	}

	close(pixelCh)
	close(resultCh)

	return img
}
