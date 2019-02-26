package element

import (
	"math"

	. "../vector"
)

type Sphere struct {
	Center Vector
	Radius float64
}

func (sphere *Sphere) Intersect(ray Ray) *HitInfo {
	oc := Subtract(ray.Origin, sphere.Center)
	b := 2.0 * Dot(ray.Direction, oc)
	c := math.Pow(oc.Length(), 2.0) - math.Pow(sphere.Radius, 2.0)
	d := b*b - 4.0*c

	t := 0.0
	intersected := true

	if 0.0 < d {
		dd := math.Sqrt(d)
		t1 := (-b + dd) / 2.0
		t2 := (-b - dd) / 2.0

		if 0.0 < t1 && 0.0 < t2 {
			t = math.Min(t1, t2)
		} else if t1 < 0.0 && t2 < 0.0 {
			intersected = false
		} else {
			t = math.Max(t1, t2)
		}
	} else if d < 0.0 {
		intersected = false
	} else {
		t = -b / 2.0
	}

	if intersected {
		pos := Add(ray.Origin, Multiply(t, ray.Direction))
		n := Normalize(Subtract(pos, sphere.Center))

		return &HitInfo{Object: sphere, Position: pos, Normal: n, T: t}
	} else {
		return nil
	}
}
