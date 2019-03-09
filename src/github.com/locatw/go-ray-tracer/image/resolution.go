package image

type Resolution struct {
	Width  int
	Height int
}

func (resolution *Resolution) Aspect() float64 {
	return float64(resolution.Width) / float64(resolution.Height)
}
