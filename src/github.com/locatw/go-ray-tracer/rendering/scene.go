package rendering

import . "github.com/locatw/go-ray-tracer/element"

type Scene struct {
	Camera Camera
	Shapes []Shape
}

func (scene *Scene) LookForIntersectedObject(ray Ray) *HitInfo {
	var minHitInfo *HitInfo

	for _, shape := range scene.Shapes {
		hitInfo := shape.Intersect(ray)

		if hitInfo == nil {
			continue
		}

		if minHitInfo == nil {
			minHitInfo = hitInfo
		} else {
			if hitInfo.T < minHitInfo.T {
				minHitInfo = hitInfo
			}
		}
	}

	return minHitInfo
}
