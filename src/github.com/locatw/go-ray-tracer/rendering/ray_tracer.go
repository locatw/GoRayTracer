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

type RayTracer struct {
	Scene            Scene
	RenderingSetting RenderingSetting
}

type RenderingSetting struct {
	Resolution                 image.Resolution
	SamplingCount              int
	TraceRecursionLimit        int
	DistanceAttenuationEnabled bool
}

func (rayTracer *RayTracer) Render() image.Image {
	rand.Seed(time.Now().UnixNano())

	camera := rayTracer.Scene.Camera

	resolution := rayTracer.RenderingSetting.Resolution
	screen := camera.CreateScreen(resolution)
	img := image.CreateImage(resolution.Width, resolution.Height)

	capacity := resolution.PixelCount()
	pixelCh := make(chan *image.Pixel, capacity)
	resultCh := make(chan *image.Pixel, capacity)

	for i := 0; i < runtime.NumCPU(); i++ {
		go rayTracer.renderPixelRoutine(&screen, pixelCh, resultCh)
	}

	for i := 0; i < len(img.Pixels); i++ {
		pixelCh <- &img.Pixels[i]
	}

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n\n", runtime.NumGoroutine())

	progressPrinter := ProgressPrinter{
		TotalCount: resolution.PixelCount(),
		Interval:   5 * resolution.Width,
		Count:      0,
	}
	finishedPixelCount := 0
	for {
		if finishedPixelCount == resolution.PixelCount() {
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

func (rayTracer *RayTracer) renderPixelRoutine(screen *Screen, pixelCh <-chan *image.Pixel, resultCh chan<- *image.Pixel) {
	for {
		pixel, ok := <-pixelCh
		if !ok {
			break
		}

		rayTracer.renderPixel(screen, pixel)

		resultCh <- pixel
	}
}

func (rayTracer *RayTracer) renderPixel(screen *Screen, pixel *image.Pixel) {
	camera := rayTracer.Scene.Camera
	setting := rayTracer.RenderingSetting

	pixelColor := image.CreateDefaultColor(image.Black)
	for _, ray := range screen.CreatePixelRays(&camera, pixel.Coordinate.X, pixel.Coordinate.Y, setting.SamplingCount) {
		color := rayTracer.traceRay(ray, setting.TraceRecursionLimit)

		pixelColor = image.AddColor(pixelColor, color)
	}

	pixel.Color = toneMap(image.DivideScalar(pixelColor, float64(setting.SamplingCount)))
}

func (rayTracer *RayTracer) traceRay(ray Ray, depth int) image.Color {
	if depth <= 0 {
		return image.CreateDefaultColor(image.Black)
	}

	hitInfo := rayTracer.Scene.LookForIntersectedObject(ray)

	if hitInfo == nil {
		return image.CreateDefaultColor(image.Black)
	}

	material := hitInfo.Object.GetMaterial()

	emissionColor := image.CreateDefaultColor(image.Black)
	if !material.Emission.NearlyEqual(emissionColor) {
		emissionColor =
			image.MultiplyScalar(Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), material.Emission)
		emissionColor = rayTracer.distanceAttenuation(ray, hitInfo, emissionColor)
	}

	diffuseColor := image.CreateDefaultColor(image.Black)
	if !material.Diffuse.NearlyEqual(diffuseColor) {
		diffuseRay := CreateDiffuseRay(ray, hitInfo)
		diffuseColor = rayTracer.traceRay(diffuseRay, depth-1)
		diffuseColor = image.MultiplyColor(material.Diffuse, diffuseColor)
		diffuseColor = rayTracer.distanceAttenuation(ray, hitInfo, diffuseColor)
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

			kr := rayTracer.reflectance(ray, normal, 1.0, *material.IndexOfRefraction)

			if kr < rand.Float64() {
				refractionColor = rayTracer.traceRay(refractRay, depth-1)
				refractionColor = rayTracer.distanceAttenuation(ray, hitInfo, refractionColor)
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
		specularColor = rayTracer.traceRay(reflectRay, depth-1)
		specularColor = image.MultiplyColor(material.Specular, specularColor)
		specularColor = rayTracer.distanceAttenuation(ray, hitInfo, specularColor)
	}

	return image.AddColorAll(emissionColor, diffuseColor, specularColor, refractionColor)
}

func (rayTracer *RayTracer) distanceAttenuation(ray Ray, hitInfo *HitInfo, color image.Color) image.Color {
	if rayTracer.RenderingSetting.DistanceAttenuationEnabled {
		v := Subtract(hitInfo.Position, ray.Origin)
		distance := v.Length()

		return image.DivideScalar(color, 1.0+0.01*math.Pow(distance, 2))
	} else {
		return color
	}
}

func (rayTracer *RayTracer) reflectance(ray Ray, normal Vector, n1 float64, n2 float64) float64 {
	// schilick's approximation
	cosTheta := Dot(Multiply(-1.0, ray.Direction), normal)
	r := math.Pow((n1-n2)/(n1+n2), 2.0)

	return r + (1.0-r)*math.Pow(1.0-cosTheta, 5)
}

func toneMap(color image.Color) image.Color {
	e := 1.0 / 2.2
	return image.Color{
		R: float32(math.Pow(float64(color.R), e)),
		G: float32(math.Pow(float64(color.G), e)),
		B: float32(math.Pow(float64(color.B), e)),
	}
}
