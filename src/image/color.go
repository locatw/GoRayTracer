package image

import "fmt"

type Color struct {
	R float32
	G float32
	B float32
}

type DefaultColor int

const (
	Black DefaultColor = iota
	White DefaultColor = iota
)

func (color DefaultColor) String() string {
	switch color {
	case Black:
		return "Black"
	case White:
		return "White"
	default:
		return "unknown"
	}
}

func CreateDefaultColor(color DefaultColor) Color {
	switch color {
	case Black:
		return Color{R: 0.0, G: 0.0, B: 0.0}
	case White:
		return Color{R: 1.0, G: 1.0, B: 1.0}
	default:
		panic(fmt.Sprintf("unknown default color : %d", color))
	}
}
