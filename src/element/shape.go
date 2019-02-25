package element

type Shape interface {
	Intersect(ray Ray) *HitInfo
}
