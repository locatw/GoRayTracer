package rendering

import "fmt"

type ProgressPrinter struct {
	TotalCount int
	Interval   int
	Count      int
}

func (printer *ProgressPrinter) Print() {
	printer.Count += 1

	if printer.Count%printer.Interval == 0 {
		progress := float64(printer.Count) / float64(printer.TotalCount) * 100.0

		fmt.Printf("Progress: %.1f %%\n", progress)
	}
}
