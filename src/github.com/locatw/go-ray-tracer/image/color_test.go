package image

import (
	"math"
	"testing"

	mathex "github.com/locatw/go-ray-tracer/math"
)

var epsilon float32 = mathex.Epsilon32()

func TestColorNearlyEqual(t *testing.T) {
	patterns := []struct {
		Color1   Color
		Color2   Color
		Expected bool
	}{
		{Color1: Color{R: 0.1, G: 0.2, B: 0.3}, Color2: Color{R: 0.1, G: 0.2, B: 0.3}, Expected: true},
		{Color1: Color{R: 0.1, G: 0.2, B: 0.3}, Color2: Color{R: 0.11, G: 0.2, B: 0.3}, Expected: false},
		{Color1: Color{R: 0.1, G: 0.2, B: 0.3}, Color2: Color{R: 0.1, G: 0.21, B: 0.3}, Expected: false},
		{Color1: Color{R: 0.1, G: 0.2, B: 0.3}, Color2: Color{R: 0.1, G: 0.2, B: 0.31}, Expected: false},
		{Color1: Color{R: 0.1, G: 0.2, B: 0.3}, Color2: Color{R: 0.1, G: 0.2, B: 0.3 + epsilon}, Expected: true},
		{Color1: Color{R: 0.1, G: 0.2, B: 0.3}, Color2: Color{R: 0.1, G: 0.2, B: 0.3 + 2.0*epsilon}, Expected: false},
	}

	for _, pattern := range patterns {
		result := pattern.Color1.NearlyEqual(pattern.Color2)
		if result != pattern.Expected {
			t.Errorf("%v.NearlyEqual(%v) must return %t, actual is %t", pattern.Color1, pattern.Color2, pattern.Expected, result)
		}
	}
}

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

func TestColorAdd(t *testing.T) {
	v1 := Color{R: 1.0, G: 2.0, B: 3.0}
	v2 := Color{R: 10.0, G: 20.0, B: 30.0}
	expected := Color{R: 11.0, G: 22.0, B: 33.0}

	result := AddColor(v1, v2)
	if result != expected {
		t.Errorf("AddColor(%v, %v) must return %v, actual is %v", v1, v2, expected, result)
	}
}

func TestAddColorAll(t *testing.T) {
	color1 := Color{R: 0.1, G: 0.2, B: 0.3}
	color2 := Color{R: 1.0, G: 2.0, B: 3.0}
	color3 := Color{R: 10.0, G: 20.0, B: 30.0}
	expected := Color{R: 11.1, G: 22.2, B: 33.3}

	result := AddColorAll(color1, color2, color3)
	if epsilon < mathex.Abs32(result.R-expected.R) ||
		epsilon < mathex.Abs32(result.G-expected.G) ||
		epsilon < mathex.Abs32(result.B-expected.B) {
		t.Errorf("AddColorAll(%v, %v, %v) must return %v, actual is %v",
			color1, color2, color3, expected, result)
	}
}

func TestMultiplyColor(t *testing.T) {
	color1 := Color{R: 0.1, G: 0.2, B: 0.3}
	color2 := Color{R: 0.01, G: 0.02, B: 0.03}
	expected := Color{R: 0.001, G: 0.004, B: 0.009}

	result := MultiplyColor(color1, color2)

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

	if epsilon < mathex.Abs32(result.R-expected.R) ||
		epsilon < mathex.Abs32(result.G-expected.G) ||
		epsilon < mathex.Abs32(result.B-expected.B) {
		t.Errorf("MultiplyScalar(%f, %v) must return %v, actual is %v", scalar, color, expected, result)
	}
}

func TestDivideScalar(t *testing.T) {
	scalar := 2.0
	color := Color{R: 0.2, G: 0.4, B: 0.6}
	expected := Color{R: 0.1, G: 0.2, B: 0.3}

	result := DivideScalar(color, scalar)

	if epsilon < mathex.Abs32(result.R-expected.R) ||
		epsilon < mathex.Abs32(result.G-expected.G) ||
		epsilon < mathex.Abs32(result.B-expected.B) {
		t.Errorf("DivideScalar(%f, %v) must return %v, actual is %v", scalar, color, expected, result)
	}
}
