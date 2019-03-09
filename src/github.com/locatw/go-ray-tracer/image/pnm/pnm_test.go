package pnm

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"

	"github.com/locatw/go-ray-tracer/image"
)

func TestWritePpm(t *testing.T) {
	img := image.CreateImage(3, 2)
	img.Data[0] = image.Color{R: 1.0, G: 0.0, B: 0.0}
	img.Data[1] = image.Color{R: 0.0, G: 1.0, B: 0.0}
	img.Data[2] = image.Color{R: 0.0, G: 0.0, B: 1.0}
	img.Data[3] = image.Color{R: 1.0, G: 1.0, B: 0.0}
	img.Data[4] = image.Color{R: 0.0, G: 1.0, B: 1.0}
	img.Data[5] = image.Color{R: 1.0, G: 0.0, B: 1.0}

	file, err := ioutil.TempFile(os.TempDir(), "image.ppm")
	if err != nil {
		t.Errorf("cannot create temp file for test")
	}

	defer file.Close()

	err = WritePpm(file.Name(), img)
	if err != nil {
		t.Errorf(err.Error())
	}

	resultFile, err := os.Open(file.Name())
	if err != nil {
		t.Errorf(err.Error())
	}

	defer resultFile.Close()

	scanner := bufio.NewScanner(resultFile)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	expectedLines := []string{
		"P3",
		"3 2",
		"255",
		"255 0 0 0 255 0 0 0 255",
		"255 255 0 0 255 255 255 0 255",
	}
	for i, expected := range expectedLines {
		if lines[i] != expected {
			t.Errorf("line %d must be \"%s\", actual is %s", i+1, expected, lines[i])
		}
	}
}
