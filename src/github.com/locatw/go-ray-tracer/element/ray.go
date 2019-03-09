package element

import (
	"math"
	"math/rand"

	mathex "github.com/locatw/go-ray-tracer/math"
	. "github.com/locatw/go-ray-tracer/vector"
)

type Ray struct {
	Origin    Vector
	Direction Vector
}

type HitInfo struct {
	Object   Shape
	Position Vector
	Normal   Vector
	T        float64
}

func CreateRay(origin Vector, direction Vector) Ray {
	return Ray{Origin: origin, Direction: Normalize(direction)}
}

func CreateDiffuseRay(ray Ray, hitInfo *HitInfo) Ray {
	// base vectors in tangent space
	u := Normalize(Cross(hitInfo.Normal, ray.Direction))
	v := Normalize(Cross(u, hitInfo.Normal))
	n := hitInfo.Normal

	r := math.Sqrt(rand.Float64())
	theta := 2.0 * math.Pi * rand.Float64()
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	z := math.Sqrt(math.Max(0.0, 1.0-x*x-y*y))

	origin := Add(hitInfo.Position, Multiply(10000.0*mathex.Epsilon(), hitInfo.Normal))
	dir := AddAll(Multiply(x, u), Multiply(y, v), Multiply(z, n))

	return CreateRay(origin, dir)
}

func CreateReflectRay(ray Ray, hitInfo *HitInfo) Ray {
	dir := Multiply(2.0*Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal), hitInfo.Normal)
	dir = Subtract(dir, Multiply(-1.0, ray.Direction))

	origin := Add(hitInfo.Position, Multiply(10000.0*mathex.Epsilon(), hitInfo.Normal))

	return CreateRay(origin, dir)
}

func CreateRefractRay(ray Ray, hitInfo *HitInfo) (Ray, bool) {
	if hitInfo.Object.GetMaterial().IndexOfRefraction == nil {
		panic("cannot create refract ray because object does not have index of refraction.")
	}

	inObject := Dot(Multiply(-1.0, ray.Direction), hitInfo.Normal) < 0.0
	iot := *hitInfo.Object.GetMaterial().IndexOfRefraction

	normal := hitInfo.Normal
	if inObject {
		normal = Multiply(-1.0, hitInfo.Normal)
	}

	eta := 1.0 / iot
	if inObject {
		eta = iot
	}

	a := Dot(Multiply(-1.0, ray.Direction), normal)
	d := 1.0 - math.Pow(eta, 2.0)*(1.0-math.Pow(a, 2.0))

	if 0.0 <= d {
		dir := Multiply(-eta, Subtract(Multiply(-1.0, ray.Direction), Multiply(a, normal)))
		dir = Subtract(dir, Multiply(math.Sqrt(d), normal))

		origin := Subtract(hitInfo.Position, Multiply(10000.0*mathex.Epsilon(), normal))

		return CreateRay(origin, dir), false
	} else {
		return Ray{}, true
	}
}
