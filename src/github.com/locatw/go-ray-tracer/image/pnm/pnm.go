package pnm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/locatw/go-ray-tracer/image"
	"github.com/locatw/go-ray-tracer/math"
)

// Convert a float value in range [0.0, 1.0] to a integer value in range [0, 255],
// and return as string value.
// If an input value is out of range, then clamp it.
func processPixelValue(value float32) string {
	return strconv.Itoa(int(math.Clamp32(math.Round32(255.0*value), 0.0, 255.0)))
}

func WritePpm(path string, image image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("P3\n")
	if err != nil {
		return err
	}

	_, err = writer.WriteString(fmt.Sprintf("%d %d\n", image.Width, image.Height))
	if err != nil {
		return err
	}

	_, err = writer.WriteString("255\n")
	if err != nil {
		return err
	}

	for h := 0; h < image.Height; h++ {
		startIndex := h * image.Width
		endIndex := (h + 1) * image.Width
		row := image.Data[startIndex:endIndex]

		rowValues := make([]string, len(row)*3)
		for i, p := range row {
			index := i * 3

			rowValues[index+0] = processPixelValue(p.R)
			rowValues[index+1] = processPixelValue(p.G)
			rowValues[index+2] = processPixelValue(p.B)
		}

		line := strings.Join(rowValues, " ") + "\n"

		_, err = writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
