package pnm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	".."
	"../../mathex"
)

// Convert a float value in range [0.0, 1.0] to a integer value in range [0, 255],
// and return as string value.
// If an input value is out of range, then clamp it.
func processPixelValue(value float32) string {
	return strconv.Itoa(int(mathex.Clamp32(mathex.Round32(255.0*value), 0.0, 255.0)))
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
		start_index := h * image.Width
		end_index := (h + 1) * image.Width
		row := image.Data[start_index:end_index]

		row_values := make([]string, len(row)*3)
		for i, p := range row {
			index := i * 3

			row_values[index+0] = processPixelValue(p.R)
			row_values[index+1] = processPixelValue(p.G)
			row_values[index+2] = processPixelValue(p.B)
		}

		line := strings.Join(row_values, " ") + "\n"

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
