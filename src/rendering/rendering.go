package rendering

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"

	. "../element"
	"../image"
	. "../vector"
)

type Coordinate struct {
	X, Y int
}

type RenderPixelResult struct {
	Coordinate Coordinate
	Color      image.Color
}

func distanceAttenuation(ray Ray, hitInfo *HitInfo, color image.Color) image.Color {
	v := Subtract(hitInfo.Position, ray.Origin)
	distance := v.Length()

	return image.DivideScalar(color, 1.0+0.01*math.Pow(distance, 2))
}

func lookForIntersectedObject(scene Scene, ray Ray) *HitInfo {
	var minHitInfo *HitInfo = nil

	for _, shape := range scene.Shapes {
		hitInfo := shape.Intersect(ray)

		if hitInfo == nil {
			continue
		}

		if minHitInfo == nil {
			minHitInfo = hitInfo
		} else {
			if hitInfo.T < minHitInfo.T {
				minHitInfo = hitInfo
			}
		}
	}

	return minHitInfo
}

func traceRay(scene Scene, ray Ray, depth int) image.Color {
	if depth <= 0 {
		return image.CreateDefaultColor(image.Black)
	}

	hitInfo := lookForIntersectedObject(scene, ray)

	if hitInfo == nil {
		return image.CreateDefaultColor(image.Black)
	}

	material := hitInfo.Object.GetMaterial()

	emission_color := image.CreateDefaultColor(image.Black)
	if !material.Emission.NearlyEqual(emission_color) {
		emission_color =
			image.MultiplyScalar(Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), material.Emission)
		emission_color = distanceAttenuation(ray, hitInfo, emission_color)
	}

	diffuse_color := image.CreateDefaultColor(image.Black)
	if !material.Diffuse.NearlyEqual(diffuse_color) {
		diffuse_ray := CreateDiffuseRay(ray, hitInfo)
		diffuse_color = traceRay(scene, diffuse_ray, depth-1)
		diffuse_color = image.MultiplyColor(material.Diffuse, diffuse_color)
		diffuse_color = distanceAttenuation(ray, hitInfo, diffuse_color)
	}

	specular_color := image.CreateDefaultColor(image.Black)
	if !material.Specular.NearlyEqual(specular_color) {
		reflect_ray := CreateReflectRay(ray, hitInfo)
		specular_color = traceRay(scene, reflect_ray, depth-1)
		specular_color = image.MultiplyColor(material.Specular, specular_color)
		specular_color = distanceAttenuation(ray, hitInfo, specular_color)
	}

	return image.AddColorAll(emission_color, diffuse_color, specular_color)
}

func createPixelRays(camera Camera, width int, height int, coord Coordinate, sampling_count int) []Ray {
	aspect := float64(width) / float64(height)

	screen_x_axis := Normalize(Cross(camera.Direction, camera.Up))
	screen_y_axis := Multiply(-1.0, camera.Up)
	screen_height := 2.0 * math.Tan(camera.Fov) / 2.0
	screen_width := screen_height * aspect

	screen_center := Add(camera.Origin, camera.Direction)

	pixel_width := screen_width / float64(width)
	pixel_height := screen_height / float64(height)
	offset := Multiply(pixel_width/2.0, screen_x_axis)
	offset = Add(offset, Multiply(pixel_height/2.0, screen_y_axis))

	left_top_pixel_center := Subtract(screen_center, Multiply(screen_width/2.0, screen_x_axis))
	left_top_pixel_center = Subtract(left_top_pixel_center, Multiply(screen_height/2.0, screen_y_axis))
	left_top_pixel_center = Add(left_top_pixel_center, offset)

	pixel_center := Add(left_top_pixel_center, Multiply(float64(coord.X)*pixel_width, screen_x_axis))
	pixel_center = Add(pixel_center, Multiply(float64(coord.Y)*pixel_height, screen_y_axis))

	rays := make([]Ray, sampling_count)
	for i := 0; i < sampling_count; i++ {
		x := rand.Float64() - 0.5
		y := rand.Float64() - 0.5

		sub_pixel_pos := AddAll(pixel_center, Multiply(x*pixel_width, screen_x_axis), Multiply(y*pixel_height, screen_y_axis))
		rays[i] = CreateRay(camera.Origin, Subtract(sub_pixel_pos, camera.Origin))
	}

	return rays
}

func renderPixel(scene Scene, width int, height int, coord Coordinate) image.Color {
	sampling_count := 10

	pixel_color := image.CreateDefaultColor(image.Black)
	for _, ray := range createPixelRays(scene.Camera, width, height, coord, sampling_count) {
		color := traceRay(scene, ray, 10)

		pixel_color = image.AddColor(pixel_color, color)
	}

	return image.DivideScalar(pixel_color, float64(sampling_count))
}

func createCoordinates(width int, height int) [][]Coordinate {
	coords := make([][]Coordinate, height)
	for i := 0; i < height; i++ {
		coords[i] = make([]Coordinate, width)

		for j := 0; j < width; j++ {
			coords[i][j].X = j
			coords[i][j].Y = i
		}
	}

	return coords
}

func renderPixelRoutine(coord_ch <-chan Coordinate, result_ch chan<- RenderPixelResult, scene Scene, width int, height int) {
	for {
		coord, ok := <-coord_ch
		if !ok {
			break
		}

		pixel_color := renderPixel(scene, width, height, coord)

		result_ch <- RenderPixelResult{Coordinate: coord, Color: pixel_color}
	}
}

func Render(scene Scene, width int, height int) image.Image {
	img := image.CreateImage(width, height)
	coords := createCoordinates(width, height)

	coord_ch := make(chan Coordinate, width*height)
	result_ch := make(chan RenderPixelResult, width*height)

	for i := 0; i < runtime.NumCPU(); i++ {
		go renderPixelRoutine(coord_ch, result_ch, scene, width, height)
	}

	for h := 0; h < len(coords); h++ {
		for w := 0; w < len(coords[h]); w++ {
			coord_ch <- coords[h][w]
		}
	}

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n\n", runtime.NumGoroutine())

	progress_printer := ProgressPrinter{TotalCount: width * height, Interval: 5 * width, Count: 0}
	finished_pixel_count := 0
	for {
		if finished_pixel_count == width*height {
			break
		}

		coord_result := <-result_ch

		index := coord_result.Coordinate.Y*img.Width + coord_result.Coordinate.X
		img.Data[index] = coord_result.Color
		finished_pixel_count += 1

		progress_printer.Print()
	}

	close(coord_ch)
	close(result_ch)

	return img
}
