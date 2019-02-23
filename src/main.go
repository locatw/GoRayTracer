package main

import (
	"fmt"
	"os"

	"./image"
	"./image/pnm"
)

func main() {
	width := 4
	height := 2

	image := image.CreateImage(width, height)

	err := pnm.WritePpm("image.ppm", image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
