package image

type Coordinate struct {
	X, Y int
}

type Pixel struct {
	Coordinate Coordinate
	Color      Color
}

type Image struct {
	Width  int
	Height int
	Pixels []Pixel
}

func CreateImage(width int, height int) Image {
	pixels := make([]Pixel, width*height)

	black := CreateDefaultColor(Black)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixels[y*width+x] = Pixel{Coordinate: Coordinate{X: x, Y: y}, Color: black}
		}
	}

	return Image{Width: width, Height: height, Pixels: pixels}
}
