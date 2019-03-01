package rendering

import (
	"math"

	. "../element"
	"../image"
	. "../vector"
)

type Coordinate struct {
	X, Y int
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

	emission_color :=
		image.MultiplyScalar(Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), material.Emission)
	emission_color = distanceAttenuation(ray, hitInfo, emission_color)

	diffuse_color :=
		image.MultiplyScalar(
			Dot(hitInfo.Normal, Multiply(-1.0, ray.Direction)), material.Diffuse)
	diffuse_color = distanceAttenuation(ray, hitInfo, diffuse_color)

	reflect_ray := CreateReflectRay(ray, hitInfo)
	reflect_color := traceRay(scene, reflect_ray, depth-1)
	specular_color := image.MultiplyColor(hitInfo.Object.GetMaterial().Specular, reflect_color)
	specular_color = distanceAttenuation(ray, hitInfo, specular_color)

	return image.AddColorAll(emission_color, diffuse_color, specular_color)
}

func createPixelRay(camera Camera, width int, height int, coord Coordinate) Ray {
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

	return CreateRay(camera.Origin, Subtract(pixel_center, camera.Origin))
}

func renderPixel(scene Scene, width int, height int, coord Coordinate) image.Color {
	ray := createPixelRay(scene.Camera, width, height, coord)

	return traceRay(scene, ray, 10)
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

func Render(scene Scene, width int, height int) image.Image {
	img := image.CreateImage(width, height)
	coords := createCoordinates(width, height)

	for h := 0; h < len(coords); h++ {
		for w := 0; w < len(coords[h]); w++ {
			img.Data[h*img.Width+w] = renderPixel(scene, width, height, coords[h][w])
		}
	}

	return img
}
