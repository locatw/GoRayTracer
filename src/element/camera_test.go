package element

import (
	"testing"

	. "../vector"
)

func TestCreateCamera(t *testing.T) {
	origin := Vector{X: 0.0, Y: 0.0, Z: 10.0}
	dir := Multiply(-2.0, CreateAxisVector(ZAxis))
	up := Multiply(2.0, CreateAxisVector(YAxis))
	fov := 60.0
	expected_dir := Multiply(-1.0, CreateAxisVector(ZAxis))
	expected_up := CreateAxisVector(YAxis)

	camera := CreateCamera(origin, dir, up, fov)

	if camera.Origin != origin {
		t.Errorf("CreateCamera(%v, %v, %v, %f) must return camera which origin is %v, actual origin is %v",
			origin, dir, up, fov, origin, camera.Origin)
	}

	if camera.Direction != expected_dir {
		t.Errorf("CreateCamera(%v, %v, %v, %f) must return camera which direction is normalized, actual direction is %v",
			origin, dir, up, fov, camera.Direction)
	}

	if camera.Up != expected_up {
		t.Errorf("CreateCamera(%v, %v, %v, %f) must return camera which up is normalized, actual up is %v",
			origin, dir, up, fov, camera.Up)
	}

	if camera.Fov != fov {
		t.Errorf("CreateCamera(%v, %v, %v, %f) must return camera which fov is %f, actual fov is %f",
			origin, dir, up, fov, fov, camera.Fov)
	}
}
