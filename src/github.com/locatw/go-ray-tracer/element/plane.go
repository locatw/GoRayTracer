package element

import . "github.com/locatw/go-ray-tracer/vector"

type Plane struct {
	Center   Vector
	Normal   Vector
	Material Material
}

func (plane *Plane) Intersect(ray Ray) *HitInfo {
	a := Dot(ray.Direction, plane.Normal)

	if a == 0.0 {
		return nil
	}

	b := Dot(Subtract(plane.Center, ray.Origin), plane.Normal)
	t := b / a
	if 0.0 < t {
		pos := Add(Multiply(t, ray.Direction), ray.Origin)

		return &HitInfo{Object: plane, Position: pos, Normal: plane.Normal, T: t}
	} else {
		return nil
	}
}

func (plane *Plane) GetMaterial() Material {
	return plane.Material
}
