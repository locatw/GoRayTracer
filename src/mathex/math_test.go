package mathex

import (
	"testing"
)

func TestRound32(t *testing.T) {
	patterns := []struct {
		value    float32
		expected float32
	}{
		{value: 1.0, expected: 1.0},
		{value: 1.4, expected: 1.0},
		{value: 1.5, expected: 2.0},
		{value: 2.0, expected: 2.0},
	}

	for _, pattern := range patterns {
		x := Round32(pattern.value)
		if x != pattern.expected {
			t.Errorf("Round32(%f) must return %f, actual %f", pattern.value, pattern.expected, x)
		}
	}
}

func TestClamp32(t *testing.T) {
	patterns := []struct {
		value    float32
		min      float32
		max      float32
		expected float32
	}{
		{value: 0.9, min: 1.0, max: 2.0, expected: 1.0},
		{value: 1.0, min: 1.0, max: 2.0, expected: 1.0},
		{value: 1.1, min: 1.0, max: 2.0, expected: 1.1},
		{value: 1.9, min: 1.0, max: 2.0, expected: 1.9},
		{value: 2.0, min: 1.0, max: 2.0, expected: 2.0},
		{value: 2.1, min: 1.0, max: 2.0, expected: 2.0},
	}

	for _, pattern := range patterns {
		x := Clamp32(pattern.value, pattern.min, pattern.max)
		if x != pattern.expected {
			t.Errorf("Clamp32(%f, %f, %f) must return %f, actual %f", pattern.value, pattern.min, pattern.max, pattern.expected, x)
		}
	}
}
