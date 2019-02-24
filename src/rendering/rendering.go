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

func traceRay(scene Scene, ray Ray) image.Color {
	return image.CreateDefaultColor(image.White)
}

func createPixelRay(camera Camera, width int, height int, coord Coordinate) Ray {
	aspect := float64(width) / float64(height)

	screen_x_axis := Cross(camera.Direction, camera.Up)
	screen_x_axis = screen_x_axis.Normalize()
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

	return traceRay(scene, ray)
}

func createCoordinates(width int, height int) [][]Coordinate {
	coords := make([][]Coordinate, height)
	for i := 0; i < height; i++ {
		coords[i] = make([]Coordinate, width)
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
