package element

import (
	"testing"

	mathex "github.com/locatw/go-ray-tracer/math"
	. "github.com/locatw/go-ray-tracer/vector"
)

var epsilon float64 = mathex.Epsilon()

func TestCreateRay(t *testing.T) {
	origin := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	dir := Vector{X: 1.0, Y: 1.0, Z: 1.0}
	expectedDir := Normalize(dir)

	ray := CreateRay(origin, dir)

	if ray.Origin != origin {
		t.Errorf("CreateRay(%v, %v) must return ray which origin is %v, actual origin is %v",
			origin, dir, origin, ray.Origin)
	}

	if ray.Direction != expectedDir {
		t.Errorf("CreateRay(%v, %v) must return ray which direction is normalized, actual direction is %v",
			origin, dir, ray.Direction)
	}
}

func TestCreateReflectRay(t *testing.T) {
	inRay := CreateRay(Vector{X: -1.0, Y: 1.0, Z: -1.0}, Vector{X: 1.0, Y: -1.0, Z: 1.0})
	plane :=
		Plane{
			Center:   CreateZeroVector(),
			Normal:   CreateAxisVector(YAxis),
			Material: CreateDefaultMaterial(),
		}
	hitInfo :=
		HitInfo{
			Object:   &plane,
			Position: CreateZeroVector(),
			Normal:   CreateAxisVector(YAxis),
			T:        1.0,
		}
	expected :=
		Ray{
			Origin:    Vector{X: 0.0, Y: 10000.0 * epsilon, Z: 0.0},
			Direction: Normalize(Vector{X: 1.0, Y: 1.0, Z: 1.0}),
		}

	reflectedRay := CreateReflectRay(inRay, &hitInfo)

	if !reflectedRay.Origin.NearlyEqual(expected.Origin) {
		t.Errorf("CreateReflectRay(%v, %v) must return reflected ray which origin is %v, actual origin is %v",
			inRay, hitInfo, expected.Origin, reflectedRay.Origin)
	}

	if !reflectedRay.Direction.NearlyEqual(expected.Direction) {
		t.Errorf("CreateReflectRay(%v, %v) must return reflected ray which direction is %v, actual direction is %v",
			inRay, hitInfo, expected.Direction, reflectedRay.Direction)
	}
}
