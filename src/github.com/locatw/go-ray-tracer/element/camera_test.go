package element

import (
	"testing"

	. "github.com/locatw/go-ray-tracer/image"
	. "github.com/locatw/go-ray-tracer/vector"
)

func TestCreateCamera(t *testing.T) {
	origin := Vector{X: 0.0, Y: 0.0, Z: 10.0}
	dir := Multiply(-2.0, CreateAxisVector(ZAxis))
	up := Multiply(2.0, CreateAxisVector(YAxis))
	resolution := Resolution{Width: 3, Height: 2}
	fov := 60.0
	expectedDir := Multiply(-1.0, CreateAxisVector(ZAxis))
	expectedUp := CreateAxisVector(YAxis)

	camera := CreateCamera(origin, dir, up, resolution, fov)

	if camera.Origin != origin {
		t.Errorf("CreateCamera(%v, %v, %v, %v, %f) must return camera which origin is %v, actual origin is %v",
			origin, dir, up, resolution, fov, origin, camera.Origin)
	}

	if camera.Direction != expectedDir {
		t.Errorf("CreateCamera(%v, %v, %v, %v, %f) must return camera which direction is normalized, actual direction is %v",
			origin, dir, up, resolution, fov, camera.Direction)
	}

	if camera.Up != expectedUp {
		t.Errorf("CreateCamera(%v, %v, %v, %v, %f) must return camera which up is normalized, actual up is %v",
			origin, dir, up, resolution, fov, camera.Up)
	}

	if camera.Resolution != resolution {
		t.Errorf("CreateCamera(%v, %v, %v, %v, %f) must return camera which resolution is %v, actual resolution is %v",
			origin, dir, up, resolution, fov, resolution, camera.Resolution)
	}

	if camera.Fov != fov {
		t.Errorf("CreateCamera(%v, %v, %v, %v, %f) must return camera which fov is %f, actual fov is %f",
			origin, dir, up, resolution, fov, fov, camera.Fov)
	}
}
