package rendering

import (
	"testing"

	. "../vector"
)

func TestCreateRay(t *testing.T) {
	origin := Vector{X: 1.0, Y: 2.0, Z: 3.0}
	dir := Vector{X: 1.0, Y: 1.0, Z: 1.0}
	expected_dir := dir.Normalize()

	ray := CreateRay(origin, dir)

	if ray.Origin != origin {
		t.Errorf("CreateRay(%v, %v) must return ray which origin is %v, actual origin is %v",
			origin, dir, origin, ray.Origin)
	}

	if ray.Direction != expected_dir {
		t.Errorf("CreateRay(%v, %v) must return ray which direction is normalized, actual direction is %v",
			origin, dir, ray.Direction)
	}
}
