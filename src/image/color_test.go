package image

import (
	"math"
	"testing"

	mathex "../math"
)

func TestDefaultColor_String(t *testing.T) {
	patterns := []struct {
		color    DefaultColor
		expected string
	}{
		{color: Black, expected: "Black"},
		{color: White, expected: "White"},
	}

	for _, pattern := range patterns {
		x := pattern.color.String()
		if x != pattern.expected {
			// cannot get string value if the target method doesn't work correctly, so pass expected value at first argument.
			t.Errorf("%s.String() must return %s, actual %s", pattern.expected, pattern.expected, x)
		}
	}

	x := DefaultColor(math.MaxInt32).String()
	if x != "unknown" {
		t.Errorf("N.String() (N is undefined constant of DefaultColor) must return %s, actual %s", "unknown", x)
	}
}

func TestCreateDefaultColor(t *testing.T) {
	patterns := []struct {
		color    DefaultColor
		expected Color
	}{
		{color: Black, expected: Color{R: 0.0, G: 0.0, B: 0.0}},
		{color: White, expected: Color{R: 1.0, G: 1.0, B: 1.0}},
	}

	for _, pattern := range patterns {
		x := CreateDefaultColor(pattern.color)
		if x.R != pattern.expected.R ||
			x.G != pattern.expected.G ||
			x.B != pattern.expected.B {

			t.Errorf("CreateDefaultColor(%s) must return Color{R: %f, G: %f, B: %f}, actual Color{R: %f, G: %f, B: %f}",
				pattern.color,
				pattern.expected.R, pattern.expected.G, pattern.expected.B,
				x.R, x.G, x.B)
		}
	}
}

func TestMultiplyColor(t *testing.T) {
	color1 := Color{R: 0.1, G: 0.2, B: 0.3}
	color2 := Color{R: 0.01, G: 0.02, B: 0.03}
	expected := Color{R: 0.001, G: 0.004, B: 0.009}

	result := MultiplyColor(color1, color2)

	epsilon := mathex.Epsilon32()
	if epsilon < mathex.Abs32(result.R-expected.R) ||
		epsilon < mathex.Abs32(result.G-expected.G) ||
		epsilon < mathex.Abs32(result.B-expected.B) {
		t.Errorf("MultiplyColor(%v, %v) must return %v, actual is %v", color1, color2, expected, result)
	}
}

func TestMultiplyScalar(t *testing.T) {
	scalar := 2.0
	color := Color{R: 0.1, G: 0.2, B: 0.3}
	expected := Color{R: 0.2, G: 0.4, B: 0.6}

	result := MultiplyScalar(scalar, color)

	epsilon := mathex.Epsilon32()
	if epsilon < mathex.Abs32(result.R-expected.R) ||
		epsilon < mathex.Abs32(result.G-expected.G) ||
		epsilon < mathex.Abs32(result.B-expected.B) {
		t.Errorf("MultiplyScalar(%f, %v) must return %v, actual is %v", scalar, color, expected, result)
	}
}
