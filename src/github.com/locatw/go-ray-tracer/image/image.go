package image

type Image struct {
	Width  int
	Height int
	Data   []Color
}

func CreateImage(width int, height int) Image {
	data := make([]Color, width*height)

	black := CreateDefaultColor(Black)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h*width+w] = black
		}
	}

	return Image{Width: width, Height: height, Data: data}
}
