package image

import "testing"

func TestCreateImage(t *testing.T) {
	width := 3
	height := 2

	image := CreateImage(width, height)

	if image.Width != width {
		t.Errorf("CreateImage(%d, %d) must return an image with Width is %d, actual width is %d",
			width, height, width, image.Width)
	}

	if image.Height != height {
		t.Errorf("CreateImage(%d, %d) must return an image with Height is %d, actual height is %d",
			width, height, height, image.Height)
	}

	expectedLength := width * height
	if len(image.Data) != expectedLength {
		t.Errorf("CreateImage(%d, %d) must return an image with length of Data is %d, actual length is %d",
			width, height, expectedLength, len(image.Data))
	}

	black := CreateDefaultColor(Black)
	isAllPixelBlack := true
	for _, data := range image.Data {
		if data != black {
			isAllPixelBlack = false
			break
		}
	}
	if !isAllPixelBlack {
		t.Errorf("CreateImage(%d, %d) must return an image with all pixels are black, but non-black pixel exists", width, height)
	}
}
